package middleware

import (
	"api_persons_echo/authorization"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Log(f func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("petición: %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

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
