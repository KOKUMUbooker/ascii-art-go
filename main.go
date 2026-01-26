package main

import ( 
"os"
"fmt"
"strings"
)

// func main() {
// 	content, _:= os.ReadFile("standard.txt")
// 	newcontent := string(content)

// 	// char := strings.Split(newcontent, "\n\n")
// 	// store := []string{}
// 	// for _, c := range char {
// 	// 	if c == "" {
// 	// 		continue
// 	// 	}
// 	// 	store = append(store, c)
// 	// }

// 	// fmt.Println(store[0])
	
// }

func main() {
	content, _ := os.ReadFile("standard.txt")
	newcontent := string(content)

	char := strings.Split(newcontent, "\n\n")
	store := []string{}
	
	for _, c := range char {
		trimmed := strings.TrimSpace(c)
		if trimmed != "" {
			store = append(store, c)
		}
	}
	fmt.Print(store[4])
}
