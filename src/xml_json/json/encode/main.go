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

func marshall(post *Post) {
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	err = ioutil.WriteFile("src/xml_json/json/marshall.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON: ", err)
		return
	}
}

func encode(post *Post) {
	jsonFile, err := os.Create("src/xml_json/json/encode.json")
	if err != nil {
		fmt.Println("Error creating JSON file: ", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding to JSON: ", err)
		return
	}
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello, World",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a nice day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	marshall(&post)
	encode(&post)

}
