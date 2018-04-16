package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tianxiaoyang2018/go-homework/bean"
	"github.com/tianxiaoyang2018/go-homework/dao"

	"github.com/gorilla/mux"
)

// curl -XGET "http://127.0.0.1:12345/users/1/relationships"

func ListRelationship(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId, _ := strconv.Atoi(params["user_id"])
	var relationships []bean.RelationshipCoreInfo = dao.ListRelationship(userId)
	json.NewEncoder(w).Encode(relationships)
}

// curl -XPUT -d '{"state":"liked"}' "http://127.0.0.1:12345/users/11231244213/relationships/21341231231"
// /users/:user_id/relationships/:other_user_id
func UpdateRelationship(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId, _ := strconv.Atoi(params["user_id"])
	otherUserId, _ := strconv.Atoi(params["other_user_id"])
	var relationshipParam bean.RelationshipCoreInfo
	_ = json.NewDecoder(req.Body).Decode(&relationshipParam)

	state := relationshipParam.State
	fmt.Println("params: user_id=", userId, ", otherUserId=", otherUserId, ",state=", state)
	dao.UpdateRelationship(userId, otherUserId, state)

	relationship := dao.GetRelationshipByUserIdAndOtherUserId(userId, otherUserId)
	json.NewEncoder(w).Encode(relationship)
}

func RegistRelationshipController() {
	var router *mux.Router = GetRouter()
	router.HandleFunc("/users/{user_id}/relationships", ListRelationship).Methods("GET")
	router.HandleFunc("/users/{user_id}/relationships/{other_user_id}", UpdateRelationship).Methods("PUT")
}
