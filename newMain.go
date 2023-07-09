package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("newFile.wiki")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf(emptyFile)
	emptyFile.Close()
}
