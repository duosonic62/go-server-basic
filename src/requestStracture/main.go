package main

import (
	"encoding/base64"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Programming Co",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Can't get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func setFlashMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No messages")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
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
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/setMessage", setFlashMessage)
	http.HandleFunc("/getMessage", showMessage)

	server.ListenAndServe()
}
