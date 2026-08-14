package ai

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/aichatmessage"
	"github.com/shuTwT/hoshikuzu/ent/aichatsession"
	"github.com/shuTwT/hoshikuzu/ent/aimodel"
	"github.com/shuTwT/hoshikuzu/ent/aiprovider"
	"github.com/shuTwT/hoshikuzu/ent/setting"
	infra_ai "github.com/shuTwT/hoshikuzu/internal/infra/ai"
	"github.com/shuTwT/hoshikuzu/internal/infra/logger"
	"github.com/shuTwT/hoshikuzu/pkg/config"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"

	openai "github.com/sashabaranov/go-openai"
)

const (
	newSessionTitle = "新对话"
	maxMessageRunes = 16000
)

var (
	ErrAIProviderNotFound      = errors.New("AI provider is not configured")
	ErrAIModelNotFound         = errors.New("no enabled AI model configured for the provider")
	ErrInvalidAIConfig         = errors.New("invalid AI provider configuration")
	ErrAIChatSessionNotFound   = errors.New("AI chat session not found")
	ErrInvalidAIChatContent    = errors.New("invalid AI chat content")
	ErrAIProviderEmptyResponse = errors.New("AI provider returned an empty response")
)

// AIService is the server-side boundary for provider configuration and chat
// data. It never returns a decrypted provider key to its callers.
type AIService interface {
	CleanupLegacySettings(ctx context.Context) error
	MigrateLegacyConfig(ctx context.Context) error
	ListProviders(ctx context.Context) ([]model.AIProviderResp, error)
	CreateProvider(ctx context.Context, req model.AIProviderReq) error
	UpdateProvider(ctx context.Context, id int, req model.AIProviderReq) error
	DeleteProvider(ctx context.Context, id int) error
	TestProvider(ctx context.Context, req model.AIProviderTestReq) error
	SyncProviderModels(ctx context.Context, id int) error
	CreateProviderModel(ctx context.Context, providerID int, req model.AIModelReq) error
	UpdateProviderModel(ctx context.Context, providerID, modelID int, req model.AIModelReq) error
	DeleteProviderModel(ctx context.Context, providerID, modelID int) error
	CreateSession(ctx context.Context, userID int) (*model.AIChatSessionResp, error)
	ListSessions(ctx context.Context, userID int) ([]model.AIChatSessionResp, error)
	ListMessages(ctx context.Context, userID, sessionID int) ([]model.AIChatMessageResp, error)
	DeleteSession(ctx context.Context, userID, sessionID int) error
	ClearSession(ctx context.Context, userID, sessionID int) error
	StreamChat(ctx context.Context, userID, sessionID int, content string, onDelta func(string) error) (*model.AIChatMessageResp, error)
	GenerateSummary(ctx context.Context, title, content string) (string, error)
}

type AIServiceImpl struct {
	client *ent.Client
	cipher func() (infra_ai.SecretCipher, error)
}

func NewAIServiceImpl(client *ent.Client) AIService {
	return &AIServiceImpl{
		client: client,
		cipher: infra_ai.NewConfigCipher,
	}
}

// CleanupLegacySettings deliberately deletes the previous, plaintext AI
// settings. They are never read or migrated into the encrypted configuration.
func (s *AIServiceImpl) CleanupLegacySettings(ctx context.Context) error {
	_, err := s.client.Setting.Delete().Where(setting.KeyIn(
		"ai",
		"openai_api_key",
		"openai_api_url",
		"openai_model",
		"openai_temperature",
		"openai_max_tokens",
		"openai_top_p",
		"openai_frequency_penalty",
		"openai_presence_penalty",
	)).Exec(ctx)
	return err
}

// MigrateLegacyConfig moves a pre-multi-provider AIConfig row (if any) into
// the AIProvider/AIModel tables so existing deployments keep their key after
// the singleton configuration was replaced. It is idempotent: it only runs
// when the legacy table exists and no provider has been created yet.
func (s *AIServiceImpl) MigrateLegacyConfig(ctx context.Context) error {
	if config.GetString(config.DATABASE_TYPE) != "sqlite" {
		return nil
	}
	// The legacy table is not part of the schema anymore, so it is read and
	// dropped through a raw connection to the same SQLite file.
	db, err := sql.Open("sqlite3", config.GetString(config.DATABASE_URL))
	if err != nil {
		return err
	}
	defer db.Close()

	var legacyTableCount int
	if err = db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='aiconfig'",
	).Scan(&legacyTableCount); err != nil {
		return err
	}
	if legacyTableCount == 0 {
		return nil
	}
	count, err := s.client.AIProvider.Query().Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	var legacy struct {
		baseURL          string
		model            string
		apiKeyCiphertext string
		temperature      float64
		maxTokens        int
		topP             float64
		frequencyPenalty float64
		presencePenalty  float64
	}
	err = db.QueryRowContext(ctx,
		`SELECT base_url, model, api_key_ciphertext, temperature, max_tokens, top_p, frequency_penalty, presence_penalty
		 FROM aiconfig WHERE config_key = 'default'`,
	).Scan(&legacy.baseURL, &legacy.model, &legacy.apiKeyCiphertext, &legacy.temperature, &legacy.maxTokens, &legacy.topP, &legacy.frequencyPenalty, &legacy.presencePenalty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.dropLegacyConfigTable(ctx, db)
		}
		return err
	}

	provider, err := s.client.AIProvider.Create().
		SetName(defaultProviderName(legacy.baseURL)).
		SetProviderType("custom").
		SetBaseURL(strings.TrimRight(legacy.baseURL, "/")).
		SetAPIKeyCiphertext(legacy.apiKeyCiphertext).
		SetTemperature(legacy.temperature).
		SetMaxTokens(legacy.maxTokens).
		SetTopP(legacy.topP).
		SetFrequencyPenalty(legacy.frequencyPenalty).
		SetPresencePenalty(legacy.presencePenalty).
		SetIsDefault(true).
		SetIsEnabled(true).
		Save(ctx)
	if err != nil {
		return err
	}
	if _, err = s.client.AIModel.Create().
		SetProviderID(provider.ID).
		SetModelName(legacy.model).
		SetIsEnabled(true).
		Save(ctx); err != nil {
		return err
	}
	if err = s.dropLegacyConfigTable(ctx, db); err != nil {
		return err
	}
	logger.Info("已将旧 AI 配置迁移为默认提供商", "provider_id", provider.ID, "model", legacy.model)
	return nil
}

func (s *AIServiceImpl) dropLegacyConfigTable(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS aiconfig"); err != nil {
		return err
	}
	return nil
}

func (s *AIServiceImpl) ListProviders(ctx context.Context) ([]model.AIProviderResp, error) {
	providers, err := s.client.AIProvider.Query().
		WithModels(func(q *ent.AIModelQuery) {
			q.Order(ent.Asc(aimodel.FieldSort), ent.Asc(aimodel.FieldID))
		}).
		Order(ent.Desc(aiprovider.FieldIsDefault), ent.Asc(aiprovider.FieldSort), ent.Asc(aiprovider.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]model.AIProviderResp, 0, len(providers))
	for _, provider := range providers {
		result = append(result, providerResponse(provider))
	}
	return result, nil
}

func (s *AIServiceImpl) CreateProvider(ctx context.Context, req model.AIProviderReq) error {
	if err := validateProvider(req); err != nil {
		return err
	}
	cipher, err := s.cipher()
	if err != nil {
		return err
	}
	apiKeyCiphertext, err := encryptAPIKey(cipher, req.APIKey)
	if err != nil {
		return err
	}
	count, err := s.client.AIProvider.Query().Count(ctx)
	if err != nil {
		return err
	}
	first := count == 0

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if req.IsDefault || first {
		if err = tx.AIProvider.Update().SetIsDefault(false).Exec(ctx); err != nil {
			return err
		}
	}
	_, err = tx.AIProvider.Create().
		SetName(strings.TrimSpace(req.Name)).
		SetProviderType(strings.TrimSpace(req.ProviderType)).
		SetBaseURL(strings.TrimRight(strings.TrimSpace(req.BaseURL), "/")).
		SetAPIKeyCiphertext(apiKeyCiphertext).
		SetTemperature(req.Temperature).
		SetMaxTokens(req.MaxTokens).
		SetTopP(req.TopP).
		SetFrequencyPenalty(req.FrequencyPenalty).
		SetPresencePenalty(req.PresencePenalty).
		SetIsDefault(req.IsDefault || first).
		SetIsEnabled(req.IsEnabled).
		SetSort(req.Sort).
		SetRemark(req.Remark).
		Save(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *AIServiceImpl) UpdateProvider(ctx context.Context, id int, req model.AIProviderReq) error {
	if err := validateProvider(req); err != nil {
		return err
	}
	provider, err := s.client.AIProvider.Query().Where(aiprovider.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrAIProviderNotFound
		}
		return err
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if req.IsDefault {
		if err = tx.AIProvider.Update().Where(aiprovider.IDNEQ(id)).SetIsDefault(false).Exec(ctx); err != nil {
			return err
		}
	}
	update := tx.AIProvider.UpdateOneID(provider.ID).
		SetName(strings.TrimSpace(req.Name)).
		SetProviderType(strings.TrimSpace(req.ProviderType)).
		SetBaseURL(strings.TrimRight(strings.TrimSpace(req.BaseURL), "/")).
		SetTemperature(req.Temperature).
		SetMaxTokens(req.MaxTokens).
		SetTopP(req.TopP).
		SetFrequencyPenalty(req.FrequencyPenalty).
		SetPresencePenalty(req.PresencePenalty).
		SetIsDefault(req.IsDefault).
		SetIsEnabled(req.IsEnabled).
		SetSort(req.Sort).
		SetRemark(req.Remark)
	if strings.TrimSpace(req.APIKey) != "" {
		cipher, err := s.cipher()
		if err != nil {
			return err
		}
		apiKeyCiphertext, err := encryptAPIKey(cipher, req.APIKey)
		if err != nil {
			return err
		}
		update = update.SetAPIKeyCiphertext(apiKeyCiphertext)
	}
	if err = update.Exec(ctx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *AIServiceImpl) DeleteProvider(ctx context.Context, id int) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.AIModel.Delete().Where(aimodel.ProviderIDEQ(id)).Exec(ctx); err != nil {
		return err
	}
	if err = tx.AIProvider.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return ErrAIProviderNotFound
		}
		return err
	}
	return tx.Commit()
}

func (s *AIServiceImpl) TestProvider(ctx context.Context, req model.AIProviderTestReq) error {
	if _, err := s.cipher(); err != nil {
		return err
	}
	if err := validateBaseURL(req.BaseURL); err != nil {
		return err
	}
	_, err := newOpenAIClient(strings.TrimSpace(req.BaseURL), strings.TrimSpace(req.APIKey)).ListModels(ctx)
	return providerError(err)
}

func (s *AIServiceImpl) SyncProviderModels(ctx context.Context, id int) error {
	provider, err := s.client.AIProvider.Query().Where(aiprovider.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrAIProviderNotFound
		}
		return err
	}
	cipher, err := s.cipher()
	if err != nil {
		return err
	}
	apiKey, err := decryptProviderKey(cipher, provider.APIKeyCiphertext)
	if err != nil {
		return err
	}
	models, err := newOpenAIClient(provider.BaseURL, apiKey).ListModels(ctx)
	if err != nil {
		return providerError(err)
	}
	if len(models.Models) == 0 {
		return ErrAIProviderEmptyResponse
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.AIModel.Delete().Where(aimodel.ProviderIDEQ(id)).Exec(ctx); err != nil {
		return err
	}
	added := 0
	for _, m := range models.Models {
		if strings.TrimSpace(m.ID) == "" {
			continue
		}
		if _, err = tx.AIModel.Create().
			SetProviderID(id).
			SetModelName(m.ID).
			SetIsEnabled(true).
			Save(ctx); err != nil {
			return err
		}
		added++
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	logger.Info("已从供应商同步模型", "provider_id", id, "provider", provider.Name, "models", added)
	return nil
}

func (s *AIServiceImpl) CreateProviderModel(ctx context.Context, providerID int, req model.AIModelReq) error {
	if err := validateModel(req); err != nil {
		return err
	}
	if _, err := s.providerForUpdate(ctx, providerID); err != nil {
		return err
	}
	_, err := s.client.AIModel.Create().
		SetProviderID(providerID).
		SetModelName(strings.TrimSpace(req.ModelName)).
		SetDisplayName(strings.TrimSpace(req.DisplayName)).
		SetIsEnabled(req.IsEnabled).
		SetSort(req.Sort).
		Save(ctx)
	return err
}

func (s *AIServiceImpl) UpdateProviderModel(ctx context.Context, providerID, modelID int, req model.AIModelReq) error {
	if err := validateModel(req); err != nil {
		return err
	}
	if _, err := s.providerForUpdate(ctx, providerID); err != nil {
		return err
	}
	m, err := s.client.AIModel.Query().
		Where(aimodel.IDEQ(modelID), aimodel.ProviderIDEQ(providerID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrAIModelNotFound
		}
		return err
	}
	_, err = s.client.AIModel.UpdateOneID(m.ID).
		SetModelName(strings.TrimSpace(req.ModelName)).
		SetDisplayName(strings.TrimSpace(req.DisplayName)).
		SetIsEnabled(req.IsEnabled).
		SetSort(req.Sort).
		Save(ctx)
	return err
}

func (s *AIServiceImpl) DeleteProviderModel(ctx context.Context, providerID, modelID int) error {
	m, err := s.client.AIModel.Query().
		Where(aimodel.IDEQ(modelID), aimodel.ProviderIDEQ(providerID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrAIModelNotFound
		}
		return err
	}
	return s.client.AIModel.DeleteOneID(m.ID).Exec(ctx)
}

func (s *AIServiceImpl) CreateSession(ctx context.Context, userID int) (*model.AIChatSessionResp, error) {
	if err := s.ensureCipher(); err != nil {
		return nil, err
	}
	session, err := s.client.AIChatSession.Create().
		SetUserID(userID).
		SetTitle(newSessionTitle).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	resp := sessionResponse(session)
	return &resp, nil
}

func (s *AIServiceImpl) ListSessions(ctx context.Context, userID int) ([]model.AIChatSessionResp, error) {
	if err := s.ensureCipher(); err != nil {
		return nil, err
	}
	sessions, err := s.client.AIChatSession.Query().
		Where(aichatsession.UserIDEQ(userID)).
		Order(ent.Desc(aichatsession.FieldUpdatedAt), ent.Desc(aichatsession.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]model.AIChatSessionResp, 0, len(sessions))
	for _, session := range sessions {
		result = append(result, sessionResponse(session))
	}
	return result, nil
}

func (s *AIServiceImpl) ListMessages(ctx context.Context, userID, sessionID int) ([]model.AIChatMessageResp, error) {
	if err := s.ensureCipher(); err != nil {
		return nil, err
	}
	if _, err := s.sessionForUser(ctx, userID, sessionID); err != nil {
		return nil, err
	}
	messages, err := s.client.AIChatMessage.Query().
		Where(aichatmessage.SessionIDEQ(sessionID)).
		Order(ent.Asc(aichatmessage.FieldCreatedAt), ent.Asc(aichatmessage.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]model.AIChatMessageResp, 0, len(messages))
	for _, message := range messages {
		result = append(result, messageResponse(message))
	}
	return result, nil
}

func (s *AIServiceImpl) DeleteSession(ctx context.Context, userID, sessionID int) error {
	if err := s.ensureCipher(); err != nil {
		return err
	}
	if _, err := s.sessionForUser(ctx, userID, sessionID); err != nil {
		return err
	}
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.AIChatMessage.Delete().Where(aichatmessage.SessionIDEQ(sessionID)).Exec(ctx); err != nil {
		return err
	}
	if err = tx.AIChatSession.DeleteOneID(sessionID).Exec(ctx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *AIServiceImpl) ClearSession(ctx context.Context, userID, sessionID int) error {
	if err := s.ensureCipher(); err != nil {
		return err
	}
	session, err := s.sessionForUser(ctx, userID, sessionID)
	if err != nil {
		return err
	}
	if _, err = s.client.AIChatMessage.Delete().Where(aichatmessage.SessionIDEQ(sessionID)).Exec(ctx); err != nil {
		return err
	}
	return s.client.AIChatSession.UpdateOneID(session.ID).SetTitle(newSessionTitle).Exec(ctx)
}

func (s *AIServiceImpl) StreamChat(ctx context.Context, userID, sessionID int, content string, onDelta func(string) error) (*model.AIChatMessageResp, error) {
	if err := s.ensureCipher(); err != nil {
		return nil, err
	}
	content = strings.TrimSpace(content)
	if content == "" || utf8.RuneCountInString(content) > maxMessageRunes {
		return nil, ErrInvalidAIChatContent
	}
	session, err := s.sessionForUser(ctx, userID, sessionID)
	if err != nil {
		return nil, err
	}

	if _, err = s.client.AIChatMessage.Create().
		SetSessionID(sessionID).
		SetRole(aichatmessage.RoleUser).
		SetContent(content).
		Save(ctx); err != nil {
		return nil, err
	}
	if session.Title == newSessionTitle {
		session.Title = makeSessionTitle(content)
	}
	if err = s.touchSession(ctx, session); err != nil {
		return nil, err
	}

	provider, err := s.providerConfig(ctx)
	if err != nil {
		return nil, err
	}
	messages, err := s.client.AIChatMessage.Query().
		Where(aichatmessage.SessionIDEQ(sessionID)).
		Order(ent.Asc(aichatmessage.FieldCreatedAt), ent.Asc(aichatmessage.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	requestMessages := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, message := range messages {
		requestMessages = append(requestMessages, openai.ChatCompletionMessage{
			Role:    string(message.Role),
			Content: message.Content,
		})
	}

	stream, err := newOpenAIClient(provider.BaseURL, provider.APIKey).CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:            provider.Model,
		Messages:         requestMessages,
		MaxTokens:        provider.MaxTokens,
		Temperature:      float32(provider.Temperature),
		TopP:             float32(provider.TopP),
		FrequencyPenalty: float32(provider.FrequencyPenalty),
		PresencePenalty:  float32(provider.PresencePenalty),
	})
	if err != nil {
		return nil, providerError(err)
	}
	defer stream.Close()

	var output strings.Builder
	for {
		response, recvErr := stream.Recv()
		if errors.Is(recvErr, io.EOF) {
			break
		}
		if recvErr != nil {
			return nil, providerError(recvErr)
		}
		if len(response.Choices) == 0 {
			continue
		}
		delta := response.Choices[0].Delta.Content
		if delta == "" {
			continue
		}
		output.WriteString(delta)
		if err := onDelta(delta); err != nil {
			return nil, err
		}
	}

	if output.Len() == 0 {
		return nil, ErrAIProviderEmptyResponse
	}
	assistant, err := s.client.AIChatMessage.Create().
		SetSessionID(sessionID).
		SetRole(aichatmessage.RoleAssistant).
		SetContent(output.String()).
		SetModel(provider.Model).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.touchSession(ctx, session); err != nil {
		return nil, err
	}
	resp := messageResponse(assistant)
	return &resp, nil
}

const (
	summarySystemPrompt   = "你是一名专业的中文文章摘要助手，请用简洁的中文总结文章内容，字数控制在 200 字以内。"
	maxSummaryInputRunes  = 8000
	maxSummaryOutputRunes = 512
)

func (s *AIServiceImpl) GenerateSummary(ctx context.Context, title, content string) (string, error) {
	if err := s.ensureCipher(); err != nil {
		return "", err
	}

	provider, err := s.providerConfig(ctx)
	if err != nil {
		return "", err
	}

	// 截断过长的文章内容，避免超出模型上下文
	input := content
	if utf8.RuneCountInString(input) > maxSummaryInputRunes {
		runes := []rune(input)
		input = string(runes[:maxSummaryInputRunes])
	}

	userContent := fmt.Sprintf("标题：%s\n\n内容：%s", title, input)

	resp, err := newOpenAIClient(provider.BaseURL, provider.APIKey).CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: provider.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: summarySystemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: userContent},
		},
		MaxTokens:        provider.MaxTokens,
		Temperature:      float32(provider.Temperature),
		TopP:             float32(provider.TopP),
		FrequencyPenalty: float32(provider.FrequencyPenalty),
		PresencePenalty:  float32(provider.PresencePenalty),
	})
	if err != nil {
		return "", providerError(err)
	}

	if len(resp.Choices) == 0 {
		return "", ErrAIProviderEmptyResponse
	}

	summary := strings.TrimSpace(resp.Choices[0].Message.Content)
	if summary == "" {
		return "", ErrAIProviderEmptyResponse
	}

	// 截断到 512 字符，匹配 post.summary 字段的 MaxLen(512)
	if utf8.RuneCountInString(summary) > maxSummaryOutputRunes {
		runes := []rune(summary)
		summary = string(runes[:maxSummaryOutputRunes])
	}

	return summary, nil
}

type providerConfig struct {
	BaseURL          string
	APIKey           string
	Model            string
	Temperature      float64
	MaxTokens        int
	TopP             float64
	FrequencyPenalty float64
	PresencePenalty  float64
}

func (s *AIServiceImpl) ensureCipher() error {
	_, err := s.cipher()
	return err
}

func (s *AIServiceImpl) providerConfig(ctx context.Context) (*providerConfig, error) {
	provider, err := s.client.AIProvider.Query().
		Where(aiprovider.IsEnabledEQ(true)).
		Order(ent.Desc(aiprovider.FieldIsDefault), ent.Asc(aiprovider.FieldSort), ent.Asc(aiprovider.FieldID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAIProviderNotFound
		}
		return nil, err
	}
	m, err := s.client.AIModel.Query().
		Where(aimodel.ProviderIDEQ(provider.ID), aimodel.IsEnabledEQ(true)).
		Order(ent.Asc(aimodel.FieldSort), ent.Asc(aimodel.FieldID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAIModelNotFound
		}
		return nil, err
	}
	cipher, err := s.cipher()
	if err != nil {
		return nil, err
	}
	apiKey, err := decryptProviderKey(cipher, provider.APIKeyCiphertext)
	if err != nil {
		return nil, err
	}
	return &providerConfig{
		BaseURL:          provider.BaseURL,
		APIKey:           apiKey,
		Model:            m.ModelName,
		Temperature:      provider.Temperature,
		MaxTokens:        provider.MaxTokens,
		TopP:             provider.TopP,
		FrequencyPenalty: provider.FrequencyPenalty,
		PresencePenalty:  provider.PresencePenalty,
	}, nil
}

func (s *AIServiceImpl) sessionForUser(ctx context.Context, userID, sessionID int) (*ent.AIChatSession, error) {
	session, err := s.client.AIChatSession.Query().
		Where(aichatsession.IDEQ(sessionID), aichatsession.UserIDEQ(userID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAIChatSessionNotFound
		}
		return nil, err
	}
	return session, nil
}

func (s *AIServiceImpl) touchSession(ctx context.Context, session *ent.AIChatSession) error {
	return s.client.AIChatSession.UpdateOneID(session.ID).SetTitle(session.Title).Exec(ctx)
}

func providerResponse(provider *ent.AIProvider) model.AIProviderResp {
	resp := model.AIProviderResp{
		ID:               provider.ID,
		Name:             provider.Name,
		ProviderType:     provider.ProviderType,
		BaseURL:          provider.BaseURL,
		APIKeyConfigured: provider.APIKeyCiphertext != "",
		Temperature:      provider.Temperature,
		MaxTokens:        provider.MaxTokens,
		TopP:             provider.TopP,
		FrequencyPenalty: provider.FrequencyPenalty,
		PresencePenalty:  provider.PresencePenalty,
		IsDefault:        provider.IsDefault,
		IsEnabled:        provider.IsEnabled,
		Sort:             provider.Sort,
		Remark:           provider.Remark,
		CreatedAt:        model.LocalTime(provider.CreatedAt),
		UpdatedAt:        model.LocalTime(provider.UpdatedAt),
		Models:           []model.AIModelResp{},
	}
	for _, m := range provider.Edges.Models {
		resp.Models = append(resp.Models, model.AIModelResp{
			ID:          m.ID,
			ModelName:   m.ModelName,
			DisplayName: m.DisplayName,
			IsEnabled:   m.IsEnabled,
			Sort:        m.Sort,
		})
	}
	return resp
}

func encryptAPIKey(cipher infra_ai.SecretCipher, apiKey string) (string, error) {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return "", nil
	}
	ciphertext, err := cipher.Encrypt(apiKey)
	if err != nil {
		return "", fmt.Errorf("encrypt AI API key: %w", err)
	}
	return ciphertext, nil
}

func decryptProviderKey(cipher infra_ai.SecretCipher, ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", nil
	}
	return cipher.Decrypt(ciphertext)
}

func (s *AIServiceImpl) providerForUpdate(ctx context.Context, providerID int) (*ent.AIProvider, error) {
	provider, err := s.client.AIProvider.Query().Where(aiprovider.IDEQ(providerID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAIProviderNotFound
		}
		return nil, err
	}
	return provider, nil
}

func defaultProviderName(rawURL string) string {
	parsed, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || parsed.Host == "" {
		return "默认提供商"
	}
	return parsed.Host
}

func sessionResponse(session *ent.AIChatSession) model.AIChatSessionResp {
	return model.AIChatSessionResp{
		ID:        session.ID,
		Title:     session.Title,
		CreatedAt: model.LocalTime(session.CreatedAt),
		UpdatedAt: model.LocalTime(session.UpdatedAt),
	}
}

func messageResponse(message *ent.AIChatMessage) model.AIChatMessageResp {
	return model.AIChatMessageResp{
		ID:        message.ID,
		Role:      string(message.Role),
		Content:   message.Content,
		Model:     message.Model,
		CreatedAt: model.LocalTime(message.CreatedAt),
	}
}

func validateProvider(req model.AIProviderReq) error {
	if strings.TrimSpace(req.Name) == "" {
		return fmt.Errorf("%w: name is required", ErrInvalidAIConfig)
	}
	if strings.TrimSpace(req.ProviderType) == "" {
		return fmt.Errorf("%w: provider_type is required", ErrInvalidAIConfig)
	}
	if err := validateBaseURL(req.BaseURL); err != nil {
		return err
	}
	if req.Temperature < 0 || req.Temperature > 2 {
		return fmt.Errorf("%w: temperature must be between 0 and 2", ErrInvalidAIConfig)
	}
	if req.MaxTokens < 1 || req.MaxTokens > 8192 {
		return fmt.Errorf("%w: max_tokens must be between 1 and 8192", ErrInvalidAIConfig)
	}
	if req.TopP < 0 || req.TopP > 1 {
		return fmt.Errorf("%w: top_p must be between 0 and 1", ErrInvalidAIConfig)
	}
	if req.FrequencyPenalty < -2 || req.FrequencyPenalty > 2 {
		return fmt.Errorf("%w: frequency_penalty must be between -2 and 2", ErrInvalidAIConfig)
	}
	if req.PresencePenalty < -2 || req.PresencePenalty > 2 {
		return fmt.Errorf("%w: presence_penalty must be between -2 and 2", ErrInvalidAIConfig)
	}
	return nil
}

func validateModel(req model.AIModelReq) error {
	if strings.TrimSpace(req.ModelName) == "" {
		return fmt.Errorf("%w: model_name is required", ErrInvalidAIConfig)
	}
	return nil
}

func validateBaseURL(rawURL string) error {
	parsed, err := url.ParseRequestURI(strings.TrimSpace(rawURL))
	if err != nil || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || (parsed.Scheme != "https" && parsed.Scheme != "http") {
		return fmt.Errorf("%w: base_url must be an absolute HTTP(S) URL", ErrInvalidAIConfig)
	}
	return nil
}

func newOpenAIClient(baseURL, apiKey string) *openai.Client {
	clientConfig := openai.DefaultConfig(apiKey)
	clientConfig.BaseURL = strings.TrimRight(baseURL, "/")
	clientConfig.HTTPClient = &http.Client{Timeout: 120 * time.Second}
	return openai.NewClientWithConfig(clientConfig)
}

func makeSessionTitle(content string) string {
	runes := []rune(content)
	if len(runes) <= 20 {
		return string(runes)
	}
	return string(runes[:20]) + "..."
}

func providerError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("AI provider request failed: %w", err)
}
