package adapters

import (
	"go-app/pkg/domain"
	"go-app/pkg/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	Repo domain.MessageRepository
	ChatGPTAPI usecase.ChatGPTAPI
	// handlerはそれぞれのインターフェイスを介して処理を実行するだけ
}

func (h *MessageHandler) SendAndReceiveMessage(c echo.Context) error {
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
