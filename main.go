package main

import (
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("garbage.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}
