package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var s []string
	rootDir := "../testingDir/"
	fileSystem := os.DirFS(rootDir)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if path == "." || strings.Contains(path, "davey") {
			fmt.Println(path + " is not required in the output, carrying on.")
		} else {
			body, err := os.ReadFile(rootDir + path)
			if err != nil {
				log.Println("unable to read file: %v", err)
			}

			s = append(s, string(body))
		}

		return nil
	})

	dt := time.Now()
	formatted := dt.Format("17052023")

	f, err := os.Create(rootDir + formatted + "-davey.log")
	if err != nil {
		log.Println("Not able to create davey.log file")
	}
	defer f.Close()

	for i := 0; i < len(s); i++ {
		_, err := f.WriteString(s[i] + "\n")
		if err != nil {
			log.Println("unable to write")
		}
		fmt.Println(s[i])
	}
}
