package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetApiKey() (string, error) {
	file, openErr := os.Open("./keylogs/api_keys.txt")
	if openErr != nil {
		return "", openErr
	}
	data, dataErr := io.ReadAll(file)
	if dataErr != nil {
		return "", dataErr
	}
	cleared := []byte{}
	for _, x := range data {
		if 33 <= int(x) && int(x) <= 126 {
			cleared = append(cleared, x)
		}
	}
	fmt.Println("Given key: " + string(cleared))
	return string(cleared), nil
}

func GetResponseFromGPT(user_query string, api_key string) (string, error) {

	url := "https://openai-hub.neuraldeep.tech/v1/chat/completions"
	payload := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": user_query},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		errMsg := fmt.Sprintf("Error marshaling JSON: %v\n", err)
		return "", errors.New(errMsg)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		errMsg := fmt.Sprintf("Error creating request: %v\n", err)
		return "", errors.New(errMsg)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error sending request: %v\n", err)
		return "", errors.New(errMsg)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error reading response body: %v\n", err)
		return "", errors.New(errMsg)
	}

	return string(body), nil
}

func main() {

	user_query := os.Args[1:]

	collect_msg := strings.Join(user_query, " ")

	api_key, loadErr := GetApiKey()

	if loadErr != nil {
		log.Println("Something went wrong during api key read")
		log.Println(loadErr)
		return
	}

	resp, respErr := GetResponseFromGPT(collect_msg, api_key)

	if respErr != nil {
		log.Println("Something went wrong during the request")
		log.Fatal(respErr)
	}

	var chatCompletion ChatCompletionResponse

	json.Unmarshal([]byte(resp), &chatCompletion)

	fmt.Println("The id of the chat completion: " + chatCompletion.Id)
	fmt.Println(chatCompletion.Choices[0].Message.Content)
	fmt.Println("Overall tokens used for answer: " + chatCompletion.Usage[0].CompletionTokens)

}
