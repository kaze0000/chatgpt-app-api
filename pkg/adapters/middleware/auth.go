package middleware

import (
	"go-app/pkg/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &usecase.JWTClaims{},
		SigningKey: []byte(secretKey),
		TokenLookup: "header:Authorization",
		AuthScheme: "Bearer",
		ErrorHandlerWithContext: JWTErrorHandler,
	}

	return middleware.JWTWithConfig(config)
}

func JWTErrorHandler(err error, c echo.Context) error {
	c.JSON(http.StatusUnauthorized, map[string]string{"error": "ログインしてください。"})
	return nil
}
