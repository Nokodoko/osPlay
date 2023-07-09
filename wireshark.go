package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("wireshark")

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
