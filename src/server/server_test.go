package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

// 共通処理
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

//共通前処理
func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/2", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content": "Updated post", "author": "Sau Sheong"}`)
	request, _ := http.NewRequest("GET", "/post/2", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
