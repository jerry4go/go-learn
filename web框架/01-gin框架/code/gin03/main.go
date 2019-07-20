package main

import "./router"

func main() {
	r := router.RouteCheck()
	r.Run()
}
