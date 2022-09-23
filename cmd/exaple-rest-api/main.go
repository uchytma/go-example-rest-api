package main

import (
	"github.com/uchytma/go-example-rest-api/api/router"
)

// @title Example REST API
// @description Example of rest API written in Golang, used for learning purposes.
// @version 1.0
// @contact.name Matěj Uchytil
// @contact.url https://github.com/uchytma/go-example-rest-api
func main() {
	router.ServeRouter()
}
