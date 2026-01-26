package main

import (
	"fmt"
	"os"
)

func getInput() string {
	if len(os.Args) < 2 {
		return ""
	}
	return os.Args[1]
}

func main() {
	input := getInput()

	if input == "" {
		return
	}

	fmt.Println(input)
}

