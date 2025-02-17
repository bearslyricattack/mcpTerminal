package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func BuildRequest(urlStr string, argsMap map[string]interface{}) (string, error) {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}
	jsonData, err := json.Marshal(argsMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal argsMap to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", baseURL.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create req: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send req: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	return string(body), nil
}
