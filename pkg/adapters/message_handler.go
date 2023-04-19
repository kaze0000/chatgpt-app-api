package adapters

import (
	"go-app/pkg/domain"
	"go-app/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	Repo domain.MessageRepository
	ChatGPTAPI usecase.ChatGPTAPI
	// handlerはそれぞれのインターフェイスを介して処理を実行するだけ
}

func (h *MessageHandler) SendMessageAndSaveResponse(c echo.Context) error {
	message := new(domain.Message)
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Repo.StoreMessage(message); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := h.ChatGPTAPI.SendMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err = h.Repo.StoreResponse(res); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *MessageHandler) GetMessagesAndResponseByUserID(c echo.Context) error {
	params := c.Param
	userID, err := strconv.Atoi(params("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なユーザーIDです"})
	}

	messages, err := h.Repo.GetMessagesByUserID(userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "メッセージの取得に失敗しました"})
	}

	// domainに切り出したほうがいい。そもそもapiからレスポンスをResponseで定義したのがミス
	type MessageWithResponse struct {
		ID int `json:"id"`
		Content string `json:"content"`
		UserID int `json:"user_id"`
		Response *domain.Response `json:"response"`
	}

	var messagesWithResponse []*MessageWithResponse

	for _, m := range messages {
		response, err := h.Repo.GetResponseByMessageID(m.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "レスポンスの取得に失敗しました"})
		}
		if response != nil {
			messageWithResponse := &MessageWithResponse{
				ID: m.ID,
				Content: m.Content,
				UserID: m.UserID,
				Response: response,
			}
			messagesWithResponse = append(messagesWithResponse, messageWithResponse)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": messagesWithResponse,
	})
}
