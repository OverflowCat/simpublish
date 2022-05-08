package api

import (
	"encoding/base64"
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
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\" + "_output"
	} else { */
	path = "_output"
	/* 	} */
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, fmt.Sprint(id)+"-") {
			return name, nil
		}
	}
	return "", errors.New("Not found")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New request!")
	passwordFromCookie := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == "password" {
			passwordFromCookie = cookie.Value
		}
	}
	if !Auth(passwordFromCookie) {
		pageURL := r.URL.String()
		// redirect to login page
		b64 := base64.StdEncoding.EncodeToString([]byte(pageURL))
		redirect_path := "/login-to-" + b64
		html := fmt.Sprintf("<html><meta http-equiv=\"refresh\" content=\"0; url=https://{%s}/\">\n", redirect_path)
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprint(w, html)
		// http.Redirect(w, r, redirect_path, http.StatusTemporaryRedirect)
		return
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
		path = "C:\\Users\\Neko\\Documents\\GitHub\\simpublish\\api\\_output\\" + filename
	} else {
		path = "_output/" + filename
	}
	log.Println("Now opening " + path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	text := string(content)
	w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Cache-Control", "s-maxage=3600") // cache
	fmt.Fprintf(w, text)
}
