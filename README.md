# Ascii Art

This Go program reads a string input from the command line and converts it into ASCII art using a predefined banner mapping stored in a file named standard.txt.
The program processes the input string, handles new lines, and prints the corresponding ASCII representation line by line.
## Key Concepts
Command Line Arguments: The program expects a single argument from the command line, which is the string to be converted.

File Handling: It opens a file to read the ASCII art representations for each character.

String Manipulation: The program processes the input string to handle new lines and convert characters to their ASCII art equivalents.

Data Structures: A map is used to associate each character (rune) with its corresponding ASCII art lines.

## How to run the program

First of all, make sure the standard.txt file is in the same directory as the main.go file.
Open the terminal and navigate to the project directory: 
```
cd ascii-art
```
Run the program and pass the text as the argument:
```
go run . "Hello"
```
The text must be written inside quotes, and you can replace it with any text you like.

The program will read standard.txt and print the input text as ASCII art in the terminal.

## Code Structure
**main()**: The entry point of the program that handles input and orchestrates the ASCII art generation.

**getBannerMapping(scanner *bufio.Scanner)**: Reads the ASCII art from the file and creates a mapping of characters to their ASCII art representations.

**getAsciiLine(s string, bannerMap map[rune][]string) string**: Prints the ASCII art for each character in the input string.

**getResultAscii(sSlice []string, bannerMap map[rune][]string)**: Processes a slice of strings and returns their combined ASCII art representation.
splitStrByNewLines(s string): Splits the input string into separate lines based on new line characters.

## Code explanation
### main():

    The program first checks if exactly one argument is provided. If not, it exits early.
    
    It then processeses the input to handle special cases like newline characters and empty strings.

    The program attempts to open standard.txt for reading. If it fails, it logs the error and terminates.

    A scanner is created to read the contents of the file, which is expected to contain the ASCII art mappings.

    The getBannerMapping function is called to create a mapping of characters to their ASCII art representations.

    After scanning, it checks for any errors that may have occurred during the scanning process.

    The input string is split into slices based on newline characters while retaining the newlines.
    
    The getResultAscii function is called to generate the final ASCII art string based on the input slices and the mapping.
    Output: Finally, the resulting ASCII art is printed to the console.

### getBannerMapping(scanner *bufio.Scanner):
    The function initializes a counter i and a variable curRune to track the current character being processed.

    It creates a bannerMap to store the ASCII representations.

    The function reads lines from the scanner, ignoring empty lines, and appends each line to the corresponding character in the bannerMap.

    It increments the curRune after every 8 lines, ensuring that the mapping is correctly aligned with the ASCII characters.

### getAsciiLine(s string, bannerMap map[rune][]string) string:
    The function initializes a strings.Builder to efficiently build the resulting ASCII line.

    It loops 8 times (as each character has 8 lines of representation) and checks if each rune in the input string exists in the bannerMap.

    If it exists, the corresponding line is appended to lineRes.

    A newline character is added after each line except the last one.

### getResultAscii(sSlice []string, bannerMap map[rune][]string):
    The function initializes a strings.Builder for the final result.

    It iterates over each string in the provided slice, converting it to runes for processing.

    If the string is a newline, it appends a newline to the result.

    Otherwise, it calls getAsciiLine to convert the string to ASCII art and appends the result.

### splitStrByNewLines(s string):
    The function initializes a slice to hold the resulting tokens and a string to accumulate characters.

    It iterates over each rune in the input string, checking for new line characters.

    When a new line is encountered, it appends the accumulated string to the tokens and resets the accumulator.

    Finally, it appends any remaining characters to the tokens.
    
## Conclusion
This Go program effectively converts a string input into ASCII art by reading character representations from a file. It demonstrates file handling, string manipulation, and the use of maps in Go.