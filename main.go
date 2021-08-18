package main

import (
	"api_persons_echo/authorization"
	"api_persons_echo/handler"
	"api_persons_echo/storage"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := authorization.LoadFiles("cmd/certificates/app.rsa", "cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("certificates couldn't be loaded: ", err)
	}
	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)
	log.Println("initialized server in 8080 port")
	err = e.Start(":8080")
	if err != nil {
		log.Println("server error: ", err)
	}
}
