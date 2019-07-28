package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/template/templates/templ.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, "Hello, World!")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/processs", process)
	server.ListenAndServe()
}
