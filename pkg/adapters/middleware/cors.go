package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware(FE_URL string) echo.MiddlewareFunc {
	fmt.Println("FE_URL", FE_URL)
	corsMiddleware := middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"*"},
		AllowOrigins: []string{FE_URL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
	    echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	})

	return corsMiddleware
}

