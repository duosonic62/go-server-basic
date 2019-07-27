package main

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h["Accept-Encoding"])
	fmt.Fprintln(w, h.Get("Accept-Encoding"))

	for _, v := range strings.Split(h.Get("Accept-Encoding"), ",") {
		fmt.Println(v)
	}
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	// Content-Length分のbyte配列を作って
	body := make([]byte, len)
	// そこにBodyを書き込み
	r.Body.Read(body)
	fmt.Println(string(body))
	fmt.Fprintln(w, string(body))
}

func form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.FormValue("hello"))
}

func multipartForm(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["upload"][0]
	//file, err := fileHeader.Open()

	file, _, err := r.FormFile("upload")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func writeBody(w http.ResponseWriter, r *http.Request) {
	str := `
{
	"book": {
		"isbn": 90729911203,
		"title": "foo",
		"author": "bar"
	}
}
`
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

func writeStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "Sorry, can't find service")
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://qiita.com/tchnkmr/items/b3d0b884db8d7d91fb1b")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func writeJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "sekky",
		Threads: []string{"1", "2", "3"},
	}

	json, _ := json2.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/form", form)
	http.HandleFunc("/multipartForm", multipartForm)
	http.HandleFunc("/writeBody", writeBody)
	http.HandleFunc("/writeStatus", writeStatus)
	http.HandleFunc("/writeHeader", writeHeader)
	http.HandleFunc("/writeJson", writeJson)

	server.ListenAndServe()
}
