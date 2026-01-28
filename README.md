# Ascii Art


This Go program reads a string input from the command line and converts it into ASCII art using a predefined banner mapping stored in a file named standard.txt.
The program processes the input string, handles new lines, and prints the corresponding ASCII representation line by line.
## Key Concepts
Command Line Arguments: The program expects a single argument from the command line, which is the string to be converted.
File Handling: It opens a file to read the ASCII art representations for each character.
String Manipulation: The program processes the input string to handle new lines and convert characters to their ASCII art equivalents.
Data Structures: A map is used to associate each character (rune) with its corresponding ASCII art lines.
## Code Structure
main(): The entry point of the program that handles input and orchestrates the ASCII art generation.
getBannerMapping(scanner *bufio.Scanner): Reads the ASCII art from the file and creates a mapping of characters to their ASCII art representations.
PrintAsciiLine(s string, bannerMap map[rune][]string): Prints the ASCII art for each character in the input string.
splitStrByNewLines(s string): Splits the input string into separate lines based on new line characters.

## Code explanation

### main():
    This function checks if the correct number of command-line arguments is provided. If not, it exits early.
    It processes the input string, replacing any \n with actual newline characters.
    The function opens the standard.txt file and initializes a scanner to read its contents.
    It calls getBannerMapping to create a mapping of runes to their ASCII representations.
    Finally, it splits the input string into lines and prints each line in ASCII format using PrintAsciiLine.

### getBannerMapping(scanner *bufio.Scanner):
    This function reads lines from the scanner and populates a map where each rune is associated with a slice of strings representing its ASCII art.
    It skips empty lines and increments a counter to track the number of lines read. After every 8 lines, it moves to the next rune.

### PrintAsciiLine(s string, bannerMap map[rune][]string):
    This function prints the ASCII representation of each character in the input string.
    It loops 8 times (as each character is represented in 8 lines) and prints the corresponding ASCII art for each rune found in the bannerMap.

### splitStrByNewLines(s string):
    This function takes a string and splits it into separate lines based on newline characters.
    It constructs a slice of strings, ensuring that each line is captured correctly, including standalone newline characters.


## Conclusion
This Go program effectively converts a string input into ASCII art by reading character representations from a file. It demonstrates file handling, string manipulation, and the use of maps in Go.