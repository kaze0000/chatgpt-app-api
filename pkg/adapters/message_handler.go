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

	var messagesWithResponse []*domain.MessageWithResponse

	for _, m := range messages {
		response, err := h.Repo.GetResponseByMessageID(m.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "レスポンスの取得に失敗しました"})
		}
		if response != nil {
			messageWithResponse := &domain.MessageWithResponse{
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

func (h *MessageHandler) UpdateMessageContent(c echo.Context) error {
	params := c.Param
	messageID, err := strconv.Atoi(params("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なメッセージIDです"})
	}

	var updatedMessage struct {
		Content string `json:"content"`
	}

	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Repo.UpdateMessageContent(messageID, updatedMessage.Content); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "メッセージの内容が更新されました"})
}
