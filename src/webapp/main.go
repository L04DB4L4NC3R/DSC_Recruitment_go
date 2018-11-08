package main

import "net/http"
import "./controller"

func main() {
	controller.Startup()
	http.ListenAndServe(":3000", nil)
}
