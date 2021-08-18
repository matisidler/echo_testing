package middleware

import (
	"api_persons_echo/authorization"
	"net/http"

	"github.com/labstack/echo"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		token := e.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			//responder token no valido
			return e.JSON(http.StatusForbidden, map[string]string{"error": "not allowed"})
		}
		return f(e)
	}
}
