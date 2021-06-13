package tests

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	if userId == 42 {
		user = User{userId, "Jack", 2}
	}

	jsonData, _ := json.Marshal(user)
	w.Write(jsonData)
}
