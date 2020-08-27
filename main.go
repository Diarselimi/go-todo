package main

import (
	"fmt"
	"net/http"

	"github.com/Diarselimi/go-todo/controllers"
)

func main() {
	controllers.RegisterControllers()
	startWebServer()
}

func startWebServer() {
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
