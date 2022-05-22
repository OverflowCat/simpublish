package api

import (
	//"encoding/base64"
	"errors"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

/* const LOCAL_DEBUG = false */

const OUTPUT_PATH = "_output"

func isDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func GetArticleById(id uint64) (string, error) {
	// read all files in the directory and find the file starting with the given ID
	files, err := ioutil.ReadDir(OUTPUT_PATH)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, fmt.Sprint(id)+"-") {
			return name, nil
		}
	}
	return "", errors.New("not found")
}

func GetArticleByTitle(title string) (string, error) {
	// read all files in the directory and find the file starting with the given ID
	files, err := ioutil.ReadDir(OUTPUT_PATH)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		name := file.Name()
		if name == title {
			return name, nil
		}

		ext := filepath.Ext(name)
		if ext != "html" && ext != "md" {
			continue
		}
		name_without_ext := strings.TrimSuffix(name, ext)
		fmt.Println("name_without_ext:", name_without_ext)
		if name_without_ext == title {
			return name, nil
		}

		id_prefix := strings.SplitN(name, "-", 2)[0]
		var name_without_id string = name
		if isDigit(id_prefix) {
			name_without_id = name[len(id_prefix)+1:]
		}
		fmt.Println("name_without_id:", name_without_id)
		name_without_ext = strings.TrimSuffix(name_without_id, ext)
		fmt.Println("name_without_ext:", name_without_ext)
		if name_without_ext == title {
			return name, nil
		}
	}
	return "", errors.New("not found")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	passwordFromCookie := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == "password" {
			passwordFromCookie = cookie.Value
		}
	}
	if !Auth(passwordFromCookie) {
		pageURL := r.URL.String()
		// b64 := base64.StdEncoding.EncodeToString([]byte(pageURL))
		encoded := url.QueryEscape(pageURL)
		redirect_path := "/login-to-" + encoded
		html := `<html>
		<meta http-equiv="refresh" content="0; url=..` + redirect_path + `" />
		</html>` // https://stackoverflow.com/a/37755644
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprint(w, html)
		// http.Redirect(w, r, redirect_path, http.StatusTemporaryRedirect)
		return
	}

	id_or_titles, ok := r.URL.Query()["id"]
	if !ok || len(id_or_titles[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}
	// Query()["id"] will return an array of items,
	// we only want the single item.
	id_or_title_str := id_or_titles[0]
	log.Println("Url Param 'id' is: " + string(id_or_title_str))
	var filename string
	var getarticleerr error
	if isDigit(id_or_title_str) {
		id, err := strconv.ParseUint(id_or_title_str, 10, 64)
		if err != nil {
			log.Println("Url Param 'id' is not a valid uint64")
			return
		}
		filename, getarticleerr = GetArticleById(id)
	} else {
		filename, getarticleerr = GetArticleByTitle(id_or_title_str)
	}
	if getarticleerr != nil {
		// 404
		html := `<html>
		<head>
		<meta charset="utf-8">
		<title>404</title>
		</head>
		<body>
		<h1>未找到文章</h1>
		<p>请检查文章编号或标题是否正确</p>
		</body>
		</html>`
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Cache-Control", "s-maxage=480")
		fmt.Fprint(w, html)
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
	if strings.HasSuffix(filename, ".md") {
		filename = filename[:len(filename)-3]
		title := strings.SplitN(filename, "-", 2)[1]

		flags := html.CommonFlags | html.CompletePage | html.HrefTargetBlank
		opts := html.RendererOptions{
			Title: title,
			Flags: flags,
		}
		renderer := html.NewRenderer(opts)
		md := []byte(text)
		text = string(markdown.ToHTML(md, nil, renderer))
	}
	w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Cache-Control", "s-maxage=3600") // cache
	fmt.Fprint(w, text)
}
