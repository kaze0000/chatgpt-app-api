package middleware

import (
	"go-app/pkg/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &usecase.JWTClaims{},
		SigningKey: []byte(secretKey),
		TokenLookup: "header:Authorization",
		AuthScheme: "Bearer",
	}

	return middleware.JWTWithConfig(config)
}
