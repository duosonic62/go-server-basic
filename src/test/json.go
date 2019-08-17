package main

import (
	"encoding/json"
	"fmt"
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

func decode(fileName string) (post Post, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
	return
}

func marshal(fileName string) (post Post, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
		return
	}

	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := decode("src/xml_json/json/post.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	_, err = marshal("src/xml_json/json/post.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
