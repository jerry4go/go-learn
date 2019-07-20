package main

import (
	"./router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
