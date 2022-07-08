package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ConstantTimeCompare(a, b string) bool {
	a_len := len(a)
	b_len := len(b)
	if a_len != b_len {
		return false
	}
	a_bytes := []byte(a)
	b_bytes := []byte(b)
	var result uint8
	for i := 0; i < a_len; i++ {
		result |= a_bytes[i] ^ b_bytes[i]
	}
	return result == 0
}

type AuthReq struct {
	Password string `json:"password"`
}

func Auth(password string) bool {
	passwd := os.Getenv("SIMPUBLISH_PASSWD")
	return ConstantTimeCompare(passwd, password)
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
