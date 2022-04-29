package api

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	// "encoding/json"
)

/* type ArticleInfo struct {
	Title string `json:"title"`
	Desc string `json:desc`	
} */

const LOCAL_DEBUG = false

func ListHandler(w http.ResponseWriter, r *http.Request) {
	/* 	fmt.Fprintf(w, "Done from another go file!\n")
	   	return */
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
	for _, file := range files {
		fmt.Fprintf(w, file.Name() + "\n")
	}
	return
}
