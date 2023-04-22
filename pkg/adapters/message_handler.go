package adapters

import (
	"go-app/pkg/domain"
	"go-app/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	mu usecase.IMessageUsecase
	// handlerはそれぞれのインターフェイスを介して処理を実行するだけ
}

func NewMessageHandler(mu usecase.IMessageUsecase) *MessageHandler {
	return &MessageHandler{mu}
}

func (h *MessageHandler) SendMessageAndSaveResponse(c echo.Context) error {
	message := new(domain.Message)
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*usecase.JWTClaims)
  userID := claims.UserID

	message.UserID = userID

	messageWithResponse, err := h.mu.SendMessageAndSaveResponse(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, messageWithResponse)
}

func (h *MessageHandler) GetMessagesAndResponseByUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*usecase.JWTClaims)
  userID := claims.UserID

	messagesWithResponse, err := h.mu.GetMessagesAndResponseByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "メッセージとレスポンスの取得に失敗しました"})
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

	if err := c.Bind(&updatedMessage); err != nil { // c.Bind: リクエストデータを検証してから、ビジネスロジックを実行
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.mu.UpdateMessageContent(messageID, updatedMessage.Content); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "メッセージの内容が更新されました"})
}

func (h *MessageHandler) DeleteMessage(c echo.Context) error {
	params := c.Param
	messageID, err := strconv.Atoi(params("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なメッセージIDです"})
	}

	if err := h.mu.DeleteMessage(messageID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "メッセージが削除されました"})
}
