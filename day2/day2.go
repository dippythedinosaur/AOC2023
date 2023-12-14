package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	ttlRed := 12
	ttlGreen := 13
	ttlBlue := 14
	intResult := 0

	fmt.Println(ttlBlue, ttlGreen, ttlRed)

	strLine1 := "Game 1: 12 blue, 15 red, 2 green; 17 red, 8 green, 5 blue; 8 red, 17 blue; 9 green, 1 blue, 4 red"

	// printFileLines("gameinput.txt")

	// finds the game number of ths line, prints the value of the game.
	strGame := FindAllByRegex(strLine1, "^.*:", "top")
	fmt.Println(strGame)
	// replaces the Game and the : with nothing
	strG := strings.Replace(strings.Replace(strGame, "Game ", "", 1), ":", "", 1)
	intG, err := strconv.Atoi(strG)
	fmt.Println(intG)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	file, err := os.Open("gameinput.txt")
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
		// fmt.Println(line)
		// get the game number
		strGame := FindAllByRegex(line, "^.*:", "top")
		// fmt.Println(strGame)
		// replaces the Game and the : with nothing
		strG := strings.Replace(strings.Replace(strGame, "Game ", "", 1), ":", "", 1)
		intG, err := strconv.Atoi(strG)
		// fmt.Println("Game:", intG)
		// get all game data
		intGreen := 0
		intRed := 0
		intBlue := 0
		var re = regexp.MustCompile(`(?m)( \d* \w*)`)
		bRes := true
		for i, match := range re.FindAllString(line, -1) {
			if i > 25 {
				fmt.Println(i)
			}

			re := regexp.MustCompile("[0-9]+")
			num := re.FindString(match)
			intnum, err := strconv.Atoi(num)

			if strings.Contains(match, "green") {
				if intnum > ttlGreen {
					bRes = false
				}
			}
			if strings.Contains(match, "red") {
				if intnum > ttlRed {
					bRes = false
				}
			}
			if strings.Contains(match, "blue") {
				if intnum > ttlBlue {
					bRes = false
				}
			}
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)

			}
		}
		if bRes != true {
			fmt.Println(bRes)
		}
		if bRes {
			intResult = intResult + intG
			fmt.Println("Game Won")
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		fmt.Println("Game: ", intG, "  Blue: ", intBlue, "  Red: ", intRed, "  Green: ", intGreen)

	}
	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result: ", intResult)
}

// This function takes a file name as input
// and prints each line of the file to the standard output
// func printFileLines(fileName string) {
// 	// Open the file
// 	file, err := os.Open(fileName)
// 	// Check for errors
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)

// 	}
// 	// Close the file when the function returns
// 	defer file.Close()
// 	// Create a scanner to read the file line by line
// 	scanner := bufio.NewScanner(file)
// 	// Loop over the scanner
// 	for scanner.Scan() {
// 		// Get the current line
// 		line := scanner.Text()
// 		// Print the line
// 		// fmt.Println(line)
// 		// get the game number
// 		strGame := FindAllByRegex(line, "^.*:", "top")
// 		// fmt.Println(strGame)
// 		// replaces the Game and the : with nothing
// 		strG := strings.Replace(strings.Replace(strGame, "Game ", "", 1), ":", "", 1)
// 		intG, err := strconv.Atoi(strG)
// 		fmt.Println("Game:", intG)
// 		// get all game data
// 		intGreen := 0
// 		intRed := 0
// 		intBlue := 0
// 		var re = regexp.MustCompile(`(?m)( \d* \w*)`)
// 		for i, match := range re.FindAllString(line, -1) {
// 			if i > 25 {
// 				fmt.Println(i)
// 			}
// 			re := regexp.MustCompile("[0-9]+")
// 			num := re.FindString(match)
// 			intnum, err := strconv.Atoi(num)

// 			if strings.Contains(match, "green") {
// 				intGreen = intGreen + intnum
// 			}
// 			if strings.Contains(match, "red") {
// 				intRed = intRed + intnum
// 			}
// 			if strings.Contains(match, "blue") {
// 				intBlue = intBlue + intnum
// 			}

// 			if err != nil {
// 				fmt.Println(err)
// 				log.Fatal(err)

// 			}

// 		}
// 		if err != nil {
// 			fmt.Println(err)
// 			log.Fatal(err)
// 		}
// 		if (intBlue <= ttlBlue) && (intRed <= ttlRed) && (intGreen <= ttlGreen) {
// 			intResult = intResult + intG
// 		}

// 		fmt.Println("Blue: ", intBlue, "  Red: ", intRed, "  Green: ", intGreen)
// 	}
// 	// Check for errors
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// }

func FindAllByRegex(s, pattern string, value string) string {
	// Compile the regex pattern
	re := regexp.MustCompile(pattern)
	// Check for errors
	// if err != nil {
	// 	return "Invalid regex pattern"
	// }
	// Replace all matches with $
	strVal := re.FindAllString(s, 1)
	return strVal[0]
}
