package main

type ReturnedMessage struct {
	Content     string   `json:"content"`
	Role        string   `json:"role"`
	Annotations []string `json:"annotations"`
}

type ReturnedChoices struct {
	FinishReason     string            `json:"finish_reason"`
	Index            string            `json:"index"`
	Message          ReturnedMessage   `json:"message"`
	ProviderSpecific map[string]string `json:"provider_specific_fields"`
}

type ReturnedUsage struct {
	CompletionTokens string `json:"completion_tokens"`
	PromptTokens     string `json:"prompt_tokens"`
	TotalTokens      string `json:"total_tokens"`
}

type ChatCompletionResponse struct {
	Id                string            `json:"id"`
	Created           int64             `json:"created"`
	Model             string            `json:"model"`
	Object            string            `json:"object"`
	SystemFingerprint string            `json:"system_fingerprint"`
	Choices           []ReturnedChoices `json:"choices"`
	Usage             []ReturnedUsage   `json:"usage"`
	ServiceTier       string            `json:"service_tier"`
}
