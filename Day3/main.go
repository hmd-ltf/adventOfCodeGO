package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type ruckSacks struct {
	firstCompartment  string
	secondCompartment string
}

func main() {
	ruckSacksDetails := readFileData()

	print(ruckSacksDetails)
}

func readFileData() []ruckSacks {
	var details []ruckSacks
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var sackDetail ruckSacks
		line := scanner.Text()
		compartmentLength, _ := findSackCompartmentLength(line)

		sackDetail.firstCompartment = line[:compartmentLength]
		sackDetail.secondCompartment = line[compartmentLength:]

		details = append(details, sackDetail)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return details
}

func findSackCompartmentLength(line string) (int, error) {
	lineLength := len(line)

	if lineLength == 0 || lineLength%2 != 0 {
		fmt.Println("The input string cannot be divided equally.")
		return 0, errors.New("The input string cannot be divided ")
	}
	return lineLength / 2, nil
}
