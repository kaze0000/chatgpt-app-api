package infra

import (
	"bytes"
	"encoding/json"
	"go-app/pkg/domain"
	"go-app/pkg/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

const openaiURL = "https://api.openai.com/v1/completions"

type chatGPTAPIClient struct {
	apiKey string
}

func NewChatGPTAPIClient(apiKey string) usecase.IChatGPTAPIClient {
	return &chatGPTAPIClient{apiKey: apiKey}
}

func (api *chatGPTAPIClient) SendMessage(message *domain.Message) (*domain.Response, error) {
	client := &http.Client{}
	data := map[string]interface{}{
		"model":      "text-davinci-003",
		"prompt":     message.Content,
		// "max_tokens": 50,
		"max_tokens": 15,
	}

	payload, err := json.Marshal(data)	// HTTPリクエストのボディに設定するデータ
																			// JSON形式でエンコードされた[]byte型のデータ
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewReader(payload)) // echo.NewRequestは第三引数にio.Reader型のデータをとる
	if err != nil {
		return nil, err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer " + api.apiKey)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// デバッグ用
	// bodyBytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return nil, err
	// }
	// fmt.Println("API response:", string(bodyBytes))

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, echo.NewHTTPError(res.StatusCode, "Chat GPT API returned an error")
	}

	resMap := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&resMap)
	if err != nil {
		return nil, err
	}

	choices := resMap["choices"].([]interface{}) //値が特定の型であることを保証する/ 失敗したらpanic
	// "choices": [
  //   {
  //     "text": "\n\nThis is indeed a test",
  //     "index": 0,
  //     "logprobs": null,
  //     "finish_reason": "length"
  //   }
  // ],
	text := choices[0].(map[string]interface{})["text"].(string)

	response := &domain.Response{MessageID: message.ID, Content: text}

	return response, nil
}
