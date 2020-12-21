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
	fmt.Println("server starting at port:5000")
	http.ListenAndServe(":5000", nil)
}
