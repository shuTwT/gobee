package model

// AIConfigUpdateReq is the full, administrator-only provider configuration.
// APIKey is write-only and is never included in a response.
type AIConfigUpdateReq struct {
	BaseURL          string  `json:"base_url"`
	APIKey           string  `json:"api_key"`
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

// AIConfigTestReq tests an unsaved provider configuration without persisting it.
type AIConfigTestReq struct {
	BaseURL string `json:"base_url"`
	APIKey  string `json:"api_key"`
}

type AIConfigResp struct {
	BaseURL          string  `json:"base_url"`
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	APIKeyConfigured bool    `json:"api_key_configured"`
}

type AIModelResp struct {
	ID string `json:"id"`
}

type AIChatSessionResp struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
}

type AIChatMessageResp struct {
	ID        int       `json:"id"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Model     string    `json:"model,omitempty"`
	CreatedAt LocalTime `json:"created_at"`
}

type AIChatStreamReq struct {
	Content string `json:"content"`
}

// AIStreamEvent is encoded as the data field of an SSE event.
type AIStreamEvent struct {
	Content   string `json:"content,omitempty"`
	MessageID int    `json:"message_id,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type AIResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
