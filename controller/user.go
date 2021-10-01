package controller

import (
	"encoding/json"
	"fmt"
	"homework-server-depl/data"
	"homework-server-depl/model"
	"homework-server-depl/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	dataString = data.DataJson
)

func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to Haeruddim API.")
}

func Follower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	username := vars["username"]

	var Data = map[string]model.Content{}

	err := json.Unmarshal([]byte(dataString), &Data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, val := range Data {
		if val.Username == username {
			res := model.ResFoll{Followers: val.Followers}

			utils.ResJson(w, http.StatusOK, res)
			return
		}
	}

	http.Error(w, "Username not found", http.StatusNotFound)
}

func DataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userid := vars["userid"]

	var Data = map[string]model.Content{}

	err := json.Unmarshal([]byte(dataString), &Data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for key, val := range Data {
		if key == userid {
			res := model.Content{Username: val.Username, Followers: val.Followers}
			utils.ResJson(w, http.StatusOK, res)
			return
		}
	}

	http.Error(w, "User Id not found", http.StatusNotFound)
}
