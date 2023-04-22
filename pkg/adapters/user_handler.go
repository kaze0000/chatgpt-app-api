// 特定のフレームワーク（Echo）を使用して、usecaseを実行できるようにする
package adapters

import (
	"go-app/pkg/domain"
	"go-app/pkg/usecase"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)
type UserHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *UserHandler {
	return &UserHandler{uu}
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := h.uu.StoreUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "新規登録に成功しました")
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(domain.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	err := godotenv.Load()
	if err != nil {
		return err
	}

	token, cookie, err := h.uu.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, token)
}
