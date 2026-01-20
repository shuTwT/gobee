package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AiApiClient struct {
	BaseUrl string
	Token   string
}

func NewAiApiClient(baseUrl string, token string) *AiApiClient {
	return &AiApiClient{BaseUrl: baseUrl, Token: token}
}

type ChatCompletionsRequestBody struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature   *float32 `json:"temperature,omitempty"`
	TopP          *float32 `json:"top_p,omitempty"`
	N             *int     `json:"n,omitempty"`
	Stream        *bool    `json:"stream,omitempty"`
	StreamOptions *struct {
		IncludeUsage *bool `json:"include_usage,omitempty"`
	} `json:"stream_options,omitempty"`
	Stop                *[]string           `json:"stop,omitempty"`
	MaxTokens           *int                `json:"max_tokens,omitempty"`
	MaxCompletionTokens *int                `json:"max_completion_tokens,omitempty"`
	PresencePenalty     *float32            `json:"presence_penalty,omitempty"`
	FrequencyPenalty    *float32            `json:"frequency_penalty,omitempty"`
	LogitBias           *map[string]float32 `json:"logit_bias,omitempty"`
	User                *string             `json:"user,omitempty"`
	Tools               *[]struct {
		Type     string `json:"type"`
		Function struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Parameters  string `json:"parameters"`
		} `json:"function"`
	} `json:"tools,omitempty"`
	ToolChoice     *string `json:"tool_choice,omitempty"`
	ResponseFormat *struct {
		Type       string          `json:"type"`
		JSONSchema *map[string]any `json:"json_schema,omitempty"`
	} `json:"response_format,omitempty"`
	Seed            *int      `json:"seed,omitempty"`
	ReasoningEffort *string   `json:"reasoning_effort,omitempty"`
	Modalities      *[]string `json:"modalities,omitempty"`
	Audio           *struct {
		Voice  string `json:"voice"`
		Format string `json:"format"`
	} `json:"audio,omitempty"`
}

// 聊天 根据对话历史创建模型响应。支持流式和非流式响应。
func (c *AiApiClient) ChatCompletions(ctx context.Context, reqBody ChatCompletionsRequestBody) error {
	url := c.BaseUrl + "/v1/chat/completions"
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	bodyData := strings.NewReader(string(jsonData))
	req, _ := http.NewRequest("POST", url, bodyData)
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(respData))

	return nil
}

type EmbeddingsRequestBody struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// 将文本转换为向量嵌入
func (c *AiApiClient) Embeddings(ctx context.Context, reqBody EmbeddingsRequestBody) error {
	url := c.BaseUrl + "/v1/embeddings"
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	bodyData := strings.NewReader(string(jsonData))
	req, _ := http.NewRequest("POST", url, bodyData)
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(respData))

	return nil
}

// 列出模型
func (c *AiApiClient) ListModels(ctx context.Context) error {
	url := c.BaseUrl + "/v1/models"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(respData))

	return nil
}

type ModerationsRequestBody struct {
	Model *string `json:"model,omitempty"`
	Input string  `json:"input"`
}

// 审查
func (c *AiApiClient) Moderations(ctx context.Context, reqBody ModerationsRequestBody) error {
	url := c.BaseUrl + "/v1/moderations"
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	bodyData := strings.NewReader(string(jsonData))
	req, _ := http.NewRequest("POST", url, bodyData)
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(respData))

	return nil
}
