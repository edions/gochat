package main

import (
	"fmt"
	"net/http"
	"gochat/api"
)

func main() {
	api.InitializeRoutes()

	fmt.Println("Server started on http://localhost:3030")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}