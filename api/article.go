package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	const LOCAL_DEBUG = false
	log.Println("New request!")
	for _, cookie := range r.Cookies() {
		log.Println("Found a cookie: ", cookie.Name, cookie.Value)
	}
	var path string
	if LOCAL_DEBUG {
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\_files\\101-测试.html"
	} else {
		path = "../output/01-测试.html"
	}
 	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	text := string(content)
    fmt.Fprintf(w, text)
}