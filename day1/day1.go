package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	str := "hddfkjdhs4sdf3"

	output := replaceByRegex(str, "[a-z]")
	fmt.Println(output)

	// data, err := readFile("datafile.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(data)

	AddCalibrationValues("datafile.txt")
}

// This function takes a string and a regex pattern as input
// and returns a new string where all matches for the pattern are replaced by $
func replaceByRegex(s, pattern string) string {
	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	// Check for errors
	if err != nil {
		return "Invalid regex pattern"
	}
	// Replace all matches with $
	return re.ReplaceAllString(s, "")
}

// This function takes a file name as input
// and prints each line of the file to the standard output
func printFileLines(fileName string) {
	// Open the file
	file, err := os.Open(fileName)
	// Check for errors
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)

	}
	// Close the file when the function returns
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Loop over the scanner
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()
		// Print the line
		fmt.Println(line)
	}
	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// This function takes a file name as input
// and prints each line of the file to the standard output
func AddCalibrationValues(fileName string) {
	// Open the file
	file, err := os.Open(fileName)
	// Check for errors
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)

	}
	// Close the file when the function returns
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Loop over the scanner
	i := 0
	for scanner.Scan() {

		// Get the current line
		line := scanner.Text()
		// Print the line
		// fmt.Println(line)
		numberStr := replaceByRegex(line, "[a-z]")
		// fmt.Println(replaceByRegex(line, "[a-z]"))

		//get the first character
		firstChar := numberStr[0:1]
		lastChar := string(numberStr[len(numberStr)-1:])
		wholeNumber := firstChar + lastChar

		fmt.Println(wholeNumber, ": ", numberStr)
		intNum, err := strconv.Atoi(wholeNumber)
		i = i + intNum

		if err != nil {
			fmt.Println(err)
		}

	}
	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}

// This function takes a file name as input
// and returns a byte slice containing the file contents
func readFile(fileName string) ([]byte, error) {
	// Use os.ReadFile to read the whole file into memory
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// Return the data and nil error
	return data, nil
}
