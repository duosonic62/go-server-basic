package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FILE_PATH = "src/data/file/"

func main() {
	data := []byte("Hello, World!\n")
	err := ioutil.WriteFile(FILE_PATH+"data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile(FILE_PATH + "data1")
	fmt.Println(string(read1))

	file1, _ := os.Create(FILE_PATH + "data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open(FILE_PATH + "data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
