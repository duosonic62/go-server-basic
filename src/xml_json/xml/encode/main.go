package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func encodeMarshal(post *Post) (err error) {
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML: ", err)
		return err
	}

	err = ioutil.WriteFile("src/xml_json/xml/marshal.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file: ", err)
		return err
	}
	return nil
}

func endoceEoncoder(post *Post) (err error) {
	xmlFile, err := os.Create("src/xml_json/xml/encoder.xml")
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding to XML: ", err)
		return
	}
	return nil
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello, World",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	err := encodeMarshal(&post)
	err = endoceEoncoder(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
}
