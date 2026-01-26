## 1. Read input from os arguments

## 2. Read banner file containing the ascii representation of characters

## 3. Create a slice that contains rune characters in the order of the banner files 

## 4. Create a function to read the input from the banner file & return an slice of string representations of the ascii characters
    - This function will go through the input line by line and isolate individual ascii string groupings that are to be appended to the result slice

[NOTE] : The length of slice in step 3 should match returned at step 4 

## 6. Create mapping for {[rune] : string} using slice in step 3 and 4

## 5. Loop over the input read from the os & for each rune in the input string, index the rune map to get what to output in the terminal