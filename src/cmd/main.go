package main

import (
	"copilot/internal/infrastructure/di"
)

func main() {
	router, _ := di.Init()
	router.Run()
}
