package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CSRFMiddleware(API_DOMAIN string) echo.MiddlewareFunc {
	csrfMiddleware := middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   "localhost", //ここで設定したドメインにのみ、ブラウザはcookieを送信する
		// CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	})

	return csrfMiddleware
}

