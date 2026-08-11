package ai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/aichatmessage"
	"github.com/shuTwT/hoshikuzu/ent/aichatsession"
	"github.com/shuTwT/hoshikuzu/ent/aiconfig"
	"github.com/shuTwT/hoshikuzu/ent/setting"
	infra_ai "github.com/shuTwT/hoshikuzu/internal/infra/ai"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"

	openai "github.com/sashabaranov/go-openai"
)

const (
	defaultAIConfigKey = "default"
	newSessionTitle    = "新对话"
	maxMessageRunes    = 16000
)

var (
	ErrAIConfigNotFound        = errors.New("AI provider is not configured")
	ErrInvalidAIConfig         = errors.New("invalid AI provider configuration")
	ErrAIChatSessionNotFound   = errors.New("AI chat session not found")
	ErrInvalidAIChatContent    = errors.New("invalid AI chat content")
	ErrAIProviderEmptyResponse = errors.New("AI provider returned an empty response")
)

// AIService is the server-side boundary for provider configuration and chat
// data. It never returns a decrypted provider key to its callers.
type AIService interface {
	CleanupLegacySettings(ctx context.Context) error
	GetConfig(ctx context.Context) (*model.AIConfigResp, error)
	SaveConfig(ctx context.Context, req model.AIConfigUpdateReq) error
	TestConfig(ctx context.Context, req model.AIConfigTestReq) error
	ListModels(ctx context.Context) ([]model.AIModelResp, error)
	CreateSession(ctx context.Context, userID int) (*model.AIChatSessionResp, error)
	ListSessions(ctx context.Context, userID int) ([]model.AIChatSessionResp, error)
	ListMessages(ctx context.Context, userID, sessionID int) ([]model.AIChatMessageResp, error)
	DeleteSession(ctx context.Context, userID, sessionID int) error
	ClearSession(ctx context.Context, userID, sessionID int) error
	StreamChat(ctx context.Context, userID, sessionID int, content string, onDelta func(string) error) (*model.AIChatMessageResp, error)
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

func (s *AIServiceImpl) GetConfig(ctx context.Context) (*model.AIConfigResp, error) {
	if _, err := s.cipher(); err != nil {
		return nil, err
	}
	config, err := s.client.AIConfig.Query().
		Where(aiconfig.ConfigKeyEQ(defaultAIConfigKey)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &model.AIConfigResp{APIKeyConfigured: false}, nil
		}
		return nil, err
	}
	cipher, err := s.cipher()
	if err != nil {
		return nil, err
	}
	if _, err := cipher.Decrypt(config.APIKeyCiphertext); err != nil {
		return nil, err
	}
	return configResponse(config), nil
}

func (s *AIServiceImpl) SaveConfig(ctx context.Context, req model.AIConfigUpdateReq) error {
	if err := validateConfig(req); err != nil {
		return err
	}
	cipher, err := s.cipher()
	if err != nil {
		return err
	}
	ciphertext, err := cipher.Encrypt(strings.TrimSpace(req.APIKey))
	if err != nil {
		return fmt.Errorf("encrypt AI API key: %w", err)
	}

	baseURL := strings.TrimRight(strings.TrimSpace(req.BaseURL), "/")
	config, err := s.client.AIConfig.Query().
		Where(aiconfig.ConfigKeyEQ(defaultAIConfigKey)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		_, err = s.client.AIConfig.Create().
			SetConfigKey(defaultAIConfigKey).
			SetBaseURL(baseURL).
			SetModel(strings.TrimSpace(req.Model)).
			SetTemperature(req.Temperature).
			SetMaxTokens(req.MaxTokens).
			SetTopP(req.TopP).
			SetFrequencyPenalty(req.FrequencyPenalty).
			SetPresencePenalty(req.PresencePenalty).
			SetAPIKeyCiphertext(ciphertext).
			Save(ctx)
		return err
	}

	return s.client.AIConfig.UpdateOneID(config.ID).
		SetBaseURL(baseURL).
		SetModel(strings.TrimSpace(req.Model)).
		SetTemperature(req.Temperature).
		SetMaxTokens(req.MaxTokens).
		SetTopP(req.TopP).
		SetFrequencyPenalty(req.FrequencyPenalty).
		SetPresencePenalty(req.PresencePenalty).
		SetAPIKeyCiphertext(ciphertext).
		Exec(ctx)
}

func (s *AIServiceImpl) TestConfig(ctx context.Context, req model.AIConfigTestReq) error {
	if _, err := s.cipher(); err != nil {
		return err
	}
	if err := validateBaseURL(req.BaseURL); err != nil {
		return err
	}
	if strings.TrimSpace(req.APIKey) == "" {
		return fmt.Errorf("%w: api_key is required", ErrInvalidAIConfig)
	}
	_, err := newOpenAIClient(strings.TrimSpace(req.BaseURL), strings.TrimSpace(req.APIKey)).ListModels(ctx)
	return providerError(err)
}

func (s *AIServiceImpl) ListModels(ctx context.Context) ([]model.AIModelResp, error) {
	provider, err := s.providerConfig(ctx)
	if err != nil {
		return nil, err
	}
	models, err := newOpenAIClient(provider.BaseURL, provider.APIKey).ListModels(ctx)
	if err != nil {
		return nil, providerError(err)
	}
	result := make([]model.AIModelResp, 0, len(models.Models))
	for _, item := range models.Models {
		if item.ID != "" {
			result = append(result, model.AIModelResp{ID: item.ID})
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result, nil
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
	cipher, err := s.cipher()
	if err != nil {
		return nil, err
	}
	config, err := s.client.AIConfig.Query().
		Where(aiconfig.ConfigKeyEQ(defaultAIConfigKey)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAIConfigNotFound
		}
		return nil, err
	}
	apiKey, err := cipher.Decrypt(config.APIKeyCiphertext)
	if err != nil {
		return nil, err
	}
	return &providerConfig{
		BaseURL:          config.BaseURL,
		APIKey:           apiKey,
		Model:            config.Model,
		Temperature:      config.Temperature,
		MaxTokens:        config.MaxTokens,
		TopP:             config.TopP,
		FrequencyPenalty: config.FrequencyPenalty,
		PresencePenalty:  config.PresencePenalty,
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

func configResponse(config *ent.AIConfig) *model.AIConfigResp {
	return &model.AIConfigResp{
		BaseURL:          config.BaseURL,
		Model:            config.Model,
		Temperature:      config.Temperature,
		MaxTokens:        config.MaxTokens,
		TopP:             config.TopP,
		FrequencyPenalty: config.FrequencyPenalty,
		PresencePenalty:  config.PresencePenalty,
		APIKeyConfigured: config.APIKeyCiphertext != "",
	}
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

func validateConfig(req model.AIConfigUpdateReq) error {
	if err := validateBaseURL(req.BaseURL); err != nil {
		return err
	}
	if strings.TrimSpace(req.APIKey) == "" {
		return fmt.Errorf("%w: api_key is required", ErrInvalidAIConfig)
	}
	if strings.TrimSpace(req.Model) == "" {
		return fmt.Errorf("%w: model is required", ErrInvalidAIConfig)
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

func validateBaseURL(rawURL string) error {
	parsed, err := url.ParseRequestURI(strings.TrimSpace(rawURL))
	if err != nil || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || (parsed.Scheme != "https" && parsed.Scheme != "http") {
		return fmt.Errorf("%w: base_url must be an absolute HTTP(S) URL", ErrInvalidAIConfig)
	}
	if !strings.HasSuffix(strings.TrimRight(parsed.Path, "/"), "/v1") {
		return fmt.Errorf("%w: base_url must include the OpenAI-compatible /v1 path", ErrInvalidAIConfig)
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
