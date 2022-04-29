package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type AuthReq struct {
	Password string `json:"password"`
}

func Auth(password string) bool {
	passwd := os.Getenv("SIMPUBLISH_PASSWD")
	return passwd == password
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var p AuthReq
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	if Auth(p.Password) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{\"result\": true}")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{\"result\": false}")
	}
}
