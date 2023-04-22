package middleware

import (
	"go-app/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckMessageOwnership(mu usecase.IMessageUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*usecase.JWTClaims)
			userID := claims.UserID

			params := c.Param
			messageID, err := strconv.Atoi(params("id"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なメッセージIDです"})
			}

			message, err := mu.GetMessageByID(messageID)
			if err != nil {
				return err
			}

			if message.UserID != userID {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "権限がありません。"})
			}

			return next(c)
		}
	}
}
