package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/core/config"
)

type RpcSendJson struct {
	ID      int64         `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type ChatGptJson struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int64  `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int64  `json:"created"`
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int64 `json:"completion_tokens"`
		PromptTokens     int64 `json:"prompt_tokens"`
		TotalTokens      int64 `json:"total_tokens"`
	} `json:"usage"`
}

func GetChatGpt(params map[string]string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	method := "POST"
	key, err := config.String("chatgpt.key")
	client := &http.Client{}
	payload := strings.NewReader(`{
		"model":"gpt-3.5-turbo",
		"messages": [{"role": "user", "content": "` + params["text"] + `"}]
	}`)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+key)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	Logger.Info(string(body))
	var data ChatGptJson
	if _err := json.Unmarshal(body, &data); _err != nil {
		Logger.Info("AppAccessToken-return-err2", _err)
	}
	return data.Choices[0].Message.Content, nil
}
