package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	if input == `\n` {
		fmt.Println()
		return
	}
	
	if input == "" {
		return
	}

	input = strings.ReplaceAll(os.Args[1],`\n`,"\n")

	// 1. Open the file
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}

	// 2. Defer closing the file until the main function returns
	defer file.Close()

	// 3. Create a new scanner object for the file
	scanner := bufio.NewScanner(file)

	// 4. Create mapping for {[rune] : ascii_string_layers_slice } as map[rune][]string
	bannerMap := getBannerMapping(scanner, file);

	// 5. Check for errors that occurred during scanning (EOF is not an error)
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error : %s", err)
	}

	sSlice := splitStrByNewLines(input)
	for _, s := range sSlice {
		sR := []rune(s)
		if len(sR) == 1 && sR[0] == '\n' {
			fmt.Println()
			continue;
		}
		PrintAsciiLine(s, bannerMap)
	}
	fmt.Println()
}

func getBannerMapping(scanner *bufio.Scanner,file *os.File) map[rune][]string {
	i := 0
	curRune := ' '
	bannerMap := map[rune][]string{
		' ': {},
	}

	// Iterate over the scanner to read line by line
	for scanner.Scan() {
		// Get the current line as a string (newline termination is stripped by default)
		line := scanner.Text()
		if len(line) == 0 { // If line is empty ignore it
			continue
		}

		bannerMap[curRune] = append(bannerMap[curRune], line)
		if i > 0 && (i+1)%8 == 0 {
			curRune++
			bannerMap[curRune] = []string{}
		}
		if len(line) > 0 {
			i++
		}

	}

	return bannerMap
}

func PrintAsciiLine(s string, bannerMap map[rune][]string) {
	for i := 0; i < 8; i++ { // Loop 8 times
		for _, r := range s {
			group, exists := bannerMap[r]
			if exists {
				fmt.Printf("%v", group[i])
			}
		}
		if i!= 7 { // Only
			fmt.Println()
		}
	}
}

func splitStrByNewLines(s string) []string {
	var tokens []string
	current := ""

	for _, r := range s {
		if r == '\n' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, "\n")
		} else {
			current += string(r)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}
