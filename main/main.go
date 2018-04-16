package main

import (
	"log"
	"net/http"

	"github.com/tianxiaoyang2018/go-homework/controller"
)

func main() {
	// regist controller
	controller.RegistTUserController()
	controller.RegistRelationshipController()
	// regis server host and port
	log.Fatal(http.ListenAndServe(":12345", controller.GetRouter()))
}
