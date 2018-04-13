package main

import "controller"
import "log"
import "net/http"

func main() {
	// regist controller
	controller.RegistTUserController()
	controller.RegistRelationshipController()
	// regis server host and port
	log.Fatal(http.ListenAndServe(":12345", controller.GetRouter()))
}
