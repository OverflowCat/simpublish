package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ArticleInfo struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Size  int64  `json:"size"`
	// Desc string `json:desc`
}

const LOCAL_DEBUG = false

func ListHandler(w http.ResponseWriter, r *http.Request) {
	passwordFromCookie := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == "password" {
			passwordFromCookie = cookie.Value
		}
	}
	if !Auth(passwordFromCookie) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600") // cache

	var path string
	path = "_output"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "[]")
		return
	}
	res := make([]ArticleInfo, len(files))
	for i, file := range files {
		log.Println(i)
		filename := file.Name()
		log.Printf("filename: %s\n", filename)
		id_str := strings.Split(filename, "-")[0]
		id, err := strconv.ParseUint(id_str, 10, 64)
		if err != nil {
			log.Panic(err)
			continue
		}
		title := strings.Join(strings.Split(filename, "-")[1:], "-")
		var filetype = "html"
		if strings.HasSuffix(title, ".html") {
			title = title[:len(title)-5]
		} else if strings.HasSuffix(title, ".md") {
			filetype = "md"
			title = title[:len(title)-3]
		}
		res[i] = ArticleInfo{Id: id, Title: title, Type: filetype, Size: file.Size()}
	}
	json.NewEncoder(w).Encode(res)
}
