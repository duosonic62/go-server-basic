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

func include(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath+"/t1.html", templatePath+"/t2.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, "Hello")
	}
}

func formatData(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func customTemplateFunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatData}
	t := template.New("customfunc.html").Funcs(funcMap)

	t, err := t.ParseFiles(templatePath + "/customfunc.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, time.Now())
	}
}

func context(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath + "/context.html")
	content := `I asked: <i>"wha't up?"</i>`
	t.Execute(w, content)
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
	http.HandleFunc("/include", include)
	http.HandleFunc("/custom", customTemplateFunc)
	http.HandleFunc("/context", context)

	server.ListenAndServe()
}
