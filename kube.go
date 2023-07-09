package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	cmd := exec.Command("kubectl", args[1])

	var buf bytes.Buffer
	cmd.Stdout = &buf

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
