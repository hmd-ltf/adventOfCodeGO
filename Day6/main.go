package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readFileData()
	fmt.Println(findMarker(input))
}

func readFileData() (input string) {

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	for scanner.Scan() {
		data := scanner.Text()
		if data != "" {
			return data
		}
	}

	return ""
}

func findMarker(input string) int {
	startInt := 0
	lastInt := 4

	for {
		duplicate := characterLastDuplicateAt(input[startInt:lastInt])
		if duplicate == 0 {
			break
		}

		startInt, lastInt = startInt+duplicate, lastInt+duplicate
	}

	return lastInt
}

func characterLastDuplicateAt(input string) int {
	var lastDuplicate int
	var existingChars []rune

Loop:
	for i := len(input) - 1; i >= 0; i-- {
		for _, val := range existingChars {
			if rune(input[i]) == val {
				lastDuplicate = i + 1
				break Loop
			}
		}

		existingChars = append(existingChars, rune(input[i]))
	}

	return lastDuplicate
}
