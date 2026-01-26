package main

import ( 
"os"
"fmt"
"strings"
)

func main() {
	content, _:= os.ReadFile("standard.txt")
	newcontent := string(content)

	char := strings.Split(newcontent, "\n\n")
	store := []string{}
	for _, c := range char {
		if c == "" {
			continue
		}
		store = append(store, c)
	}
	str := os.Args[1]
	fmt.Print(Display(str, store))		
}

// func Display(str string, arstr []string) string {
// 	var result strings.Builder
// 	charLines := make([][]string, len(str))
// 	for i, c := range str {
// 		index := int(c) - 32
// 		if index >= 0 && index < len(arstr) {
// 			charLines[i] = strings.Split(arstr[index], "\n")
// 		}
// 	}
// 	if len(charLines) == 0 || len(charLines[0]) == 0 {
// 		return ""
// 	}
// 	height := len(charLines[0])

// 	for row := 0; row < height; row++ {
// 		for col := 0; col < len(charLines); col++ {
// 			if row < len(charLines[col]) {
// 				result.WriteString(charLines[col][row])
// 			}
// 		}
// 		result.WriteString("\n")
// 	}

// 	return result.String()
// }

func Display(str string, arstr []string) string {
    var result strings.Builder

    runes := []rune(str)
    charLines := make([][]string, len(runes))

    const charHeight = 8 // standard ASCII-art height

    for i, c := range runes {
        index := int(c) - 32

        if index >= 0 && index < len(arstr) {
            charLines[i] = strings.Split(arstr[index], "\n")
        } else {
            // fill with empty space if invalid character
            charLines[i] = make([]string, charHeight)
            for j := 0; j < charHeight; j++ {
                charLines[i][j] = "       "
            }
        }
    }

    for row := 0; row < charHeight; row++ {
        for col := 0; col < len(charLines); col++ {
            result.WriteString(charLines[col][row])
        }
        result.WriteString("\n")
    }

    return result.String()
}