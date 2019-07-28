package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

const templatePath = "src/template/templates"

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath + "/templ.html")
	rand.Seed(time.Now().Unix())
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, rand.Intn(10) > 5)
	}
}

func iterate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath + "/iterate.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thi", "Fri", "Sat", "Sun"}
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, daysOfWeek)
	}
}

func with(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath + "/with.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, "Hello")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/processs", process)
	http.HandleFunc("/iterate", iterate)
	http.HandleFunc("/with", with)

	server.ListenAndServe()
}
