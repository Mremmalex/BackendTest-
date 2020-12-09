package main

import (
	"backEndTest/model"
	"backEndTest/routes"
	"fmt"
	"net/http"
)

func main() {
	model.Init()
	routes.Init()
	fmt.Println("server starting at port:8080")
	http.ListenAndServe(":5000", nil)
}

