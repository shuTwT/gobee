package model

// AIProviderReq is the administrator-only provider payload used by both
// create and update. APIKey is write-only: an empty value keeps the stored
// key untouched on update and stores no key on create (local services such
// as Ollama need no key).
type AIProviderReq struct {
	Name             string  `json:"name" validate:"required,max=100"`
	ProviderType     string  `json:"provider_type" validate:"required,max=50"`
	BaseURL          string  `json:"base_url" validate:"required"`
	APIKey           string  `json:"api_key"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	IsDefault        bool    `json:"is_default"`
	IsEnabled        bool    `json:"is_enabled"`
	Sort             int     `json:"sort"`
	Remark           string  `json:"remark"`
}

// AIProviderTestReq tests an unsaved provider connection without persisting it.
// APIKey may be empty for keyless local services.
type AIProviderTestReq struct {
	ProviderType string `json:"provider_type"`
	BaseURL      string `json:"base_url" validate:"required"`
	APIKey       string `json:"api_key"`
}

// AIProviderResp is the provider view returned to the admin UI. The API key
// itself is never exposed, only whether one has been configured.
type AIProviderResp struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	ProviderType     string        `json:"provider_type"`
	BaseURL          string        `json:"base_url"`
	APIKeyConfigured bool          `json:"api_key_configured"`
	Temperature      float64       `json:"temperature"`
	MaxTokens        int           `json:"max_tokens"`
	TopP             float64       `json:"top_p"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	PresencePenalty  float64       `json:"presence_penalty"`
	IsDefault        bool          `json:"is_default"`
	IsEnabled        bool          `json:"is_enabled"`
	Sort             int           `json:"sort"`
	Remark           string        `json:"remark"`
	CreatedAt        LocalTime     `json:"created_at"`
	UpdatedAt        LocalTime     `json:"updated_at"`
	Models           []AIModelResp `json:"models"`
}

// AIModelReq is the payload for a model entry under a provider.
type AIModelReq struct {
	ModelName   string `json:"model_name" validate:"required,max=255"`
	DisplayName string `json:"display_name"`
	IsEnabled   bool   `json:"is_enabled"`
	Sort        int    `json:"sort"`
}

type AIModelResp struct {
	ID          int    `json:"id"`
	ModelName   string `json:"model_name"`
	DisplayName string `json:"display_name,omitempty"`
	IsEnabled   bool   `json:"is_enabled"`
	Sort        int    `json:"sort"`
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
