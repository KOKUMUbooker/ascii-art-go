package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

const helloNoNewline = ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `

const helloSingleNewline = ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

`

const helloDoubleNewline = ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       `

type TestGroup struct {
	name     string
	input    string
	expected string
}

func TestMain(t *testing.T) {
	tests := []TestGroup{
		{
			name:     "String with no newline",
			input:    "hello",
			expected: helloNoNewline,
		},
		{
			name:     "String with single newline",
			input:    "Hello\n",
			expected: helloSingleNewline,
		},
		{
			name:     "String with double newline",
			input:    "Hello\n\nThere",
			expected: helloDoubleNewline,
		},
	}

	bannerMap, err := prepareTestBannerMap()
	if err != nil {
		t.Fatal(err)
	}

	for i, tableTest := range tests {
		tt := tableTest
		t.Run(tt.name, func(t *testing.T) {
			sSlice := splitStrByNewLines(tt.input)
			res := getResultAscii(sSlice, bannerMap)
			if res != tt.expected {
				// t.Errorf("got %q, want %q", res, tt.expected)
				t.Errorf("Test %v failed ", i+1)
			}
		})
	}
}

func prepareTestBannerMap() (map[rune][]string, error) {
	file, err := os.Open("standard.txt")
	if err != nil {
		err := fmt.Sprintf("failed to open file: %s", err)
		return nil, errors.New(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	bannerMap := getBannerMapping(scanner)

	return bannerMap, nil
}
