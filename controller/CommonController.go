package controller

import (
	"sync"

	"github.com/gorilla/mux"
)

var router *mux.Router

var once sync.Once

func GetRouter() *mux.Router {
	once.Do(func() {
		router = mux.NewRouter()
	})
	return router
}
