package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	resp, err := http.Get("https://en.wikipedia.org/wiki/Dow_Jones_Industrial_Average")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf(resp.Status)
	}
	defer resp.Body.Close()
	f, err := os.Create("tickers.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	out, err := exec.Command("cat", "tickers.html").Output()
	if err != nil {
		fmt.Printf("file contents:\n, %s", out)
	}
}
