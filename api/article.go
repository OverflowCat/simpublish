package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/* const LOCAL_DEBUG = false */

func GetArticleById(id uint64) (string, error) {
	// Read all files in the directory and find the file starting with the given ID
	var path string
/* 	if LOCAL_DEBUG {
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\" + "_files"
	} else { */
		path = "_files"
/* 	} */
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, fmt.Sprint(id) + "-") {
			return name, nil
		}
	}
	return "", errors.New("Not found")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New request!")
	for _, cookie := range r.Cookies() {
		log.Println("Found a cookie: ", cookie.Name, cookie.Value)
	}

	ids, ok := r.URL.Query()["id"]
    if !ok || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }
    // Query()["id"] will return an array of items, 
    // we only want the single item.
    id_str := ids[0]
    log.Println("Url Param 'id' is: " + string(id_str))
	id, err := strconv.ParseUint(id_str, 10, 64)
	if err != nil {
		log.Panicln("Incorrect article ID")
		return
	}

	filename, err := GetArticleById(id)
	if err != nil {
		log.Panic(err)
		return
	}

	var path string
	if LOCAL_DEBUG {
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\_files\\" + filename
	} else {
		path = "_files/" + filename
	}
	log.Println("Now opening " + path)
 	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	text := string(content)
    fmt.Fprintf(w, text)
}