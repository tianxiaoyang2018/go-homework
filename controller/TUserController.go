package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tianxiaoyang2018/go-homework/bean"
	"github.com/tianxiaoyang2018/go-homework/dao"

	"github.com/gorilla/mux"
)

func ListTUser(w http.ResponseWriter, req *http.Request) {
	var userList []bean.UserCoreInfo = dao.ListUser()
	json.NewEncoder(w).Encode(userList)
}

func GetTUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	var user bean.UserCoreInfo = dao.GetTUser(id)
	json.NewEncoder(w).Encode(user)
}

func InserTUser(w http.ResponseWriter, req *http.Request) {
	var user bean.UserCoreInfo
	_ = json.NewDecoder(req.Body).Decode(&user)
	dao.InsertUser(user)
	user = dao.GetTUserByName(user.Name)
	json.NewEncoder(w).Encode(user)
}

func RegistTUserController() {
	var router *mux.Router = GetRouter()
	router.HandleFunc("/users", ListTUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetTUser).Methods("GET")
	router.HandleFunc("/users", InserTUser).Methods("POST")
}
