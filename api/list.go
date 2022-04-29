package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ArticleInfo struct {
	Id uint64 `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	Size int64 `json:"size"`
	// Desc string `json:desc`
}

const LOCAL_DEBUG = true

func ListHandler(w http.ResponseWriter, r *http.Request) {
	var path string
	if LOCAL_DEBUG {
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\" + "_files"
	} else {
		path = "_files"
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panic(err)
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
		if strings.HasSuffix(title, ".html") {
			title = title[:len(title)-5]
		}
		res[i] = ArticleInfo{Id: id, Title: title, Type: "html", Size: file.Size()}
	}
	log.Printf("res: %v\n", res)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
