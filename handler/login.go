package handler

import (
	"api_persons_echo/authorization"
	"api_persons_echo/model"
	"net/http"

	"github.com/labstack/echo"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(e echo.Context) error {
	data := model.Login{}
	//usamos e.Bind para volcar la info del body en data.
	err := e.Bind(&data)
	if err != nil {
		response := newResponse(Error, "struct isn't correct", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	if !isLoginValid(&data) {
		response := newResponse(Error, "user or password is not valid", nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "couldn't generate token", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "ok", dataToken)
	return e.JSON(http.StatusOK, response)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@edteam.com" && data.Password == "123456"
}
