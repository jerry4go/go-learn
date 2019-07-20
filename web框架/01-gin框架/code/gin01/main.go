package main

import (
	"./initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	router.Run()
}
