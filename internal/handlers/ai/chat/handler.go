package chat

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/user"
	infra_ai "github.com/shuTwT/hoshikuzu/internal/infra/ai"
	"github.com/shuTwT/hoshikuzu/internal/middleware"
	ai_service "github.com/shuTwT/hoshikuzu/internal/services/ai/chat"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
	"github.com/valyala/fasthttp"
)

type AIHandler struct {
	service ai_service.AIService
	client  *ent.Client
}

func NewAIHandler(service ai_service.AIService, client *ent.Client) *AIHandler {
	return &AIHandler{service: service, client: client}
}

// @Summary 获取 AI 配置
// @Description 获取脱敏后的 OpenAI 兼容配置，仅超级管理员可用
// @Tags 后台管理接口/AI
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=model.AIConfigResp}
// @Failure 403 {object} model.HttpError
// @Failure 503 {object} model.HttpError
// @Router /api/v1/ai/config [get]
func (h *AIHandler) GetConfig(c *fiber.Ctx) error {
	if err := h.requireSuperAdmin(c); err != nil {
		return err
	}
	config, err := h.service.GetConfig(c.Context())
	if err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", config))
}

// @Summary 保存 AI 配置
// @Description 保存完整的 OpenAI 兼容配置，API Key 只写入服务端密文
// @Tags 后台管理接口/AI
// @Accept json
// @Produce json
// @Param req body model.AIConfigUpdateReq true "AI 配置"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 403 {object} model.HttpError
// @Failure 503 {object} model.HttpError
// @Router /api/v1/ai/config [put]
func (h *AIHandler) SaveConfig(c *fiber.Ctx) error {
	if err := h.requireSuperAdmin(c); err != nil {
		return err
	}
	var req model.AIConfigUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	if err := h.service.SaveConfig(c.Context(), req); err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 测试 AI 配置
// @Description 服务端测试未保存的 OpenAI 兼容配置
// @Tags 后台管理接口/AI
// @Accept json
// @Produce json
// @Param req body model.AIConfigTestReq true "测试配置"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 403 {object} model.HttpError
// @Failure 502 {object} model.HttpError
// @Router /api/v1/ai/config/test [post]
func (h *AIHandler) TestConfig(c *fiber.Ctx) error {
	if err := h.requireSuperAdmin(c); err != nil {
		return err
	}
	var req model.AIConfigTestReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	if err := h.service.TestConfig(c.Context(), req); err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 获取 AI 模型列表
// @Description 由服务端使用已保存配置请求供应商模型列表
// @Tags 后台管理接口/AI
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]model.AIModelResp}
// @Failure 403 {object} model.HttpError
// @Failure 502 {object} model.HttpError
// @Router /api/v1/ai/config/models [get]
func (h *AIHandler) ListModels(c *fiber.Ctx) error {
	if err := h.requireSuperAdmin(c); err != nil {
		return err
	}
	models, err := h.service.ListModels(c.Context())
	if err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", models))
}

// @Summary 获取聊天会话列表
// @Description 获取当前登录用户自己的聊天会话
// @Tags 后台管理接口/AI聊天
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]model.AIChatSessionResp}
// @Failure 401 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions [get]
func (h *AIHandler) ListSessions(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	sessions, err := h.service.ListSessions(c.Context(), userID)
	if err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", sessions))
}

// @Summary 创建聊天会话
// @Description 为当前登录用户创建一个空聊天会话
// @Tags 后台管理接口/AI聊天
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=model.AIChatSessionResp}
// @Failure 401 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions [post]
func (h *AIHandler) CreateSession(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	session, err := h.service.CreateSession(c.Context(), userID)
	if err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", session))
}

// @Summary 获取会话消息
// @Description 获取当前用户会话中的全部消息
// @Tags 后台管理接口/AI聊天
// @Produce json
// @Param id path int true "会话 ID"
// @Success 200 {object} model.HttpSuccess{data=[]model.AIChatMessageResp}
// @Failure 404 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions/{id}/messages [get]
func (h *AIHandler) ListMessages(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	sessionID, err := parseID(c)
	if err != nil {
		return err
	}
	messages, err := h.service.ListMessages(c.Context(), userID, sessionID)
	if err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", messages))
}

// @Summary 删除聊天会话
// @Description 删除当前用户会话及其全部消息
// @Tags 后台管理接口/AI聊天
// @Produce json
// @Param id path int true "会话 ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 404 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions/{id} [delete]
func (h *AIHandler) DeleteSession(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	sessionID, err := parseID(c)
	if err != nil {
		return err
	}
	if err := h.service.DeleteSession(c.Context(), userID, sessionID); err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 清空聊天会话
// @Description 删除当前会话的全部消息并重置标题
// @Tags 后台管理接口/AI聊天
// @Produce json
// @Param id path int true "会话 ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 404 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions/{id}/messages [delete]
func (h *AIHandler) ClearSession(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	sessionID, err := parseID(c)
	if err != nil {
		return err
	}
	if err := h.service.ClearSession(c.Context(), userID, sessionID); err != nil {
		return h.writeError(c, err)
	}
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 流式发送聊天消息
// @Description 将当前会话全部历史提交给 OpenAI 兼容模型并通过 SSE 返回增量内容
// @Tags 后台管理接口/AI聊天
// @Accept json
// @Produce text/event-stream
// @Param id path int true "会话 ID"
// @Param req body model.AIChatStreamReq true "聊天消息"
// @Success 200 {string} string
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 503 {object} model.HttpError
// @Router /api/v1/ai/chat/sessions/{id}/stream [post]
func (h *AIHandler) Stream(c *fiber.Ctx) error {
	userID, err := currentUserID(c)
	if err != nil {
		return err
	}
	sessionID, err := parseID(c)
	if err != nil {
		return err
	}
	var req model.AIChatStreamReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	c.Set(fiber.HeaderContentType, "text/event-stream; charset=utf-8")
	c.Set(fiber.HeaderCacheControl, "no-cache, no-transform")
	c.Set(fiber.HeaderConnection, "keep-alive")
	c.Set("X-Accel-Buffering", "no")
	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(writer *bufio.Writer) {
		send := func(event string, payload model.AIStreamEvent) error {
			data, marshalErr := json.Marshal(payload)
			if marshalErr != nil {
				return marshalErr
			}
			if _, writeErr := fmt.Fprintf(writer, "event: %s\ndata: %s\n\n", event, data); writeErr != nil {
				return writeErr
			}
			return writer.Flush()
		}

		assistant, streamErr := h.service.StreamChat(c.Context(), userID, sessionID, req.Content, func(delta string) error {
			return send("delta", model.AIStreamEvent{Content: delta})
		})
		if streamErr != nil {
			_ = send("error", model.AIStreamEvent{Code: streamErrorCode(streamErr), Message: streamErrorMessage(streamErr)})
			return
		}
		_ = send("done", model.AIStreamEvent{MessageID: assistant.ID})
	}))
	return nil
}

func (h *AIHandler) requireSuperAdmin(c *fiber.Ctx) error {
	loginUser, err := currentUserID(c)
	if err != nil {
		return err
	}
	u, err := h.client.User.Query().Where(user.IDEQ(loginUser)).WithRole().Only(c.Context())
	if err != nil {
		return h.writeError(c, err)
	}
	if u.Edges.Role == nil || u.Edges.Role.Code != "superAdmin" {
		return c.Status(fiber.StatusForbidden).JSON(model.NewError(fiber.StatusForbidden, "AI configuration requires super administrator role"))
	}
	return nil
}

func currentUserID(c *fiber.Ctx) (int, error) {
	loginUser := middleware.GetCurrentUser(c)
	if loginUser == nil {
		return 0, c.Status(fiber.StatusUnauthorized).JSON(model.NewError(fiber.StatusUnauthorized, "Authentication required"))
	}
	return loginUser.ID, nil
}

func parseID(c *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return 0, c.Status(fiber.StatusBadRequest).JSON(model.NewError(fiber.StatusBadRequest, "Invalid session ID"))
	}
	return id, nil
}

func (h *AIHandler) writeError(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := "AI request failed"
	switch {
	case errors.Is(err, ai_service.ErrInvalidAIConfig), errors.Is(err, ai_service.ErrInvalidAIChatContent):
		status = fiber.StatusBadRequest
		message = err.Error()
	case errors.Is(err, ai_service.ErrAIChatSessionNotFound):
		status = fiber.StatusNotFound
		message = "AI chat session not found"
	case errors.Is(err, ai_service.ErrAIConfigNotFound), errors.Is(err, infra_ai.ErrConfigEncryptionKeyUnavailable), errors.Is(err, infra_ai.ErrConfigDecryptionFailed):
		status = fiber.StatusServiceUnavailable
		message = "AI service is not configured"
	case errors.Is(err, ai_service.ErrAIProviderEmptyResponse):
		status = fiber.StatusBadGateway
		message = "AI provider returned an empty response"
	default:
		if stringsContainsProviderError(err) {
			status = fiber.StatusBadGateway
			message = "AI provider request failed"
		}
	}
	return c.Status(status).JSON(model.NewError(status, message))
}

func streamErrorCode(err error) string {
	switch {
	case errors.Is(err, ai_service.ErrAIConfigNotFound), errors.Is(err, infra_ai.ErrConfigEncryptionKeyUnavailable), errors.Is(err, infra_ai.ErrConfigDecryptionFailed):
		return "configuration_unavailable"
	case errors.Is(err, ai_service.ErrAIChatSessionNotFound):
		return "session_not_found"
	case errors.Is(err, ai_service.ErrInvalidAIChatContent):
		return "invalid_content"
	case errors.Is(err, ai_service.ErrAIProviderEmptyResponse):
		return "empty_provider_response"
	default:
		return "provider_error"
	}
}

func streamErrorMessage(err error) string {
	switch {
	case errors.Is(err, ai_service.ErrAIConfigNotFound), errors.Is(err, infra_ai.ErrConfigEncryptionKeyUnavailable), errors.Is(err, infra_ai.ErrConfigDecryptionFailed):
		return "AI service is not configured"
	case errors.Is(err, ai_service.ErrAIChatSessionNotFound):
		return "AI chat session not found"
	case errors.Is(err, ai_service.ErrInvalidAIChatContent):
		return err.Error()
	case errors.Is(err, ai_service.ErrAIProviderEmptyResponse):
		return "AI provider returned an empty response"
	default:
		return "AI provider request failed"
	}
}

func stringsContainsProviderError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "AI provider request failed")
}
