package controller

import "github.com/gorilla/mux"

var router *mux.Router

func GetRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	return router
}
