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

func GetResponseFromGPT(user_query string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY environment variable not set.")
	}

	url := "https://api.openai.com/v1/chat/completions"
	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
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
	req.Header.Set("Authorization", "Bearer "+apiKey)

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

	resp, respErr := GetResponseFromGPT(collect_msg)

	if respErr != nil {
		log.Println("Something went wrong during the request")
		log.Fatal(resp)
	}

	fmt.Println(resp)

}
