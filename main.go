package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createFile()
	writeFile()
	readFile()
}

func createFile() {
	var file, err = os.Create("./files/file.txt")
	defer file.Close() //Close file created
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Created Successfully")
}

func writeFile() {
	var err = ioutil.WriteFile("./files/file.txt", []byte("Insert text here"), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Updated Successfully")
}

func readFile() {
	var content, err = ioutil.ReadFile("./files/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Content:", string(content))
}
