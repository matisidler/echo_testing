package handler

import (
	"api_persons_echo/middleware"

	"github.com/labstack/echo"
)

//esta funcion routePerson me permite crear todos los handlers de person (create, update, etc)
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	persons := e.Group("/persons")
	persons.Use(middleware.Authentication)
	persons.POST("", h.create)
	persons.DELETE("/:id", h.delete)
	persons.GET("/:id", h.getById)
	persons.PUT("/:id", h.update)
	persons.GET("", h.getAll)

}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
