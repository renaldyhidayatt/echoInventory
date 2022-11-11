package middleware

import (
	"echoinventory/schemas"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/api/login",
	"/api/register",
}

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "JWT Token is missing"
}

func WebSecurityConfig(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &schemas.JWtMetaRequest{},
		SigningKey: []byte("secret"),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	path := e.Path()
	for _, p := range whiteListPaths {
		if path == p {
			return true
		}
	}
	return false
}
