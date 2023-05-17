package main

import (
	"fmt"
	"io/fs"
	"os"
	"log"
)

func main() {
	fileSystem := os.DirFS("../testingDir")
	
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}


		fmt.Println(path)

		body, err := os.ReadFile("../testingDir/" + path)
		if err != nil {
			log.Println("unable to read file: %v", err)
		}

		fmt.Println(string(body))
		return nil
		})
}