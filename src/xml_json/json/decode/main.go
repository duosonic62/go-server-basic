package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decoder(jsonFile *os.File) {
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON: ", err)
		}
		fmt.Println(post)
	}
}

func marshal(jsonFile *os.File) {
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}

func main() {
	jsonFile, err := os.Open("src/xml_json/json/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	marshal(jsonFile)
	decoder(jsonFile)
}
