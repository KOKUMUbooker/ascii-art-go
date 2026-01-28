package main

import (
	"bufio"
	"fmt"
)

func getBannerMapping(scanner *bufio.Scanner) map[rune][]string {
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
			if i <= 126 {
				bannerMap[curRune] = []string{}
			}
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
		if i != 7 { // Only
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
