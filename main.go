package main

import (
	"api_persons_echo/authorization"
	"api_persons_echo/handler"
	"api_persons_echo/storage"
	"log"

	"github.com/labstack/echo"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("certificates couldn't be loaded: ", err)
	}
	store := storage.NewMemory()

	e := echo.New()

	//	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)
	log.Println("initialized server in 8080 port")
	err = e.Start(":8080")
	if err != nil {
		log.Println("server error: ", err)
	}
}
