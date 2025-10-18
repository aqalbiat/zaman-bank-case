package main

type ReturnedChoices struct {
}

type ChatCompletionResponse struct {
	Id                string            `json:"id"`
	Created           int64             `json:"created"`
	Model             string            `json:"model"`
	Object            string            `json:"object"`
	SystemFingerprint string            `json:"system_fingerprint"`
	Choices           []ReturnedChoices `json:"choices"`
}
