package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// for i := ' '; i <= 126; i++ {
	// 	fmt.Printf("%v = %v \n", i, string(i))
	// }

	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	if input == "" {
		fmt.Println("")
	}

	// 1. Open the file
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}

	// 2. Defer closing the file until the main function returns
	defer file.Close()

	// 3. Create a new scanner object for the file
	scanner := bufio.NewScanner(file)

	i := 0
	curRune := ' '
	res := map[rune][]string{
		' ': {},
	}
	// 4. Iterate over the scanner to read line by line
	for scanner.Scan() {
		// Get the current line as a string (newline termination is stripped by default)
		line := scanner.Text()
		if len(line) == 0 { // If line is empty ignore it
			continue
		}

		// fmt.Println(line, "len : ", len(line), " i : ", i, "is divisible by 8 : ", i%8 == 0)

		res[curRune] = append(res[curRune], line)
		if i > 0 && (i+1)%8 == 0 {
			curRune++
			res[curRune] = []string{}
		}
		if len(line) > 0 {
			i++
		}

	}
	// fmt.Println("\nNumber of lines : ", i)
	// fmt.Println("res items : ", len(res))
	// for key, val := range res {
	// 	fmt.Println("Key : ", string(key))
	// 	for _, s := range val {
	// 		fmt.Println(s)
	// 	}
	// 	fmt.Println()
	// }

	// 5. Check for errors that occurred during scanning (EOF is not an error)
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error: %s", err)
	}

	for i := 0; i < 8; i++ { // Loop 8 times
		for _, r := range input {
			group, exists := res[r]
			if exists {
				fmt.Printf("%v", group[i])
			}
		}
		fmt.Println()
	}
}
