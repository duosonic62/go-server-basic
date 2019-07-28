package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"golang.org/x/net/http2"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func worldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

// Chain Handler Log
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	http.HandleFunc("/hello_func", log(helloHandler))
	http.HandleFunc("/world_func", log(worldHandler))

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}
