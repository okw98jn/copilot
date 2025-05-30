package main

import (
	"copilot/internal/adapter/controller"
	"copilot/internal/infrastructure/router"
)

func main() {
	router := router.New(controller.NewUserController())
	router.Run()
}
