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
	ruckSacksDetails, _ := readFileData()

	print(ruckSacksDetails)
}

func readFileData() ([]ruckSacks, error) {
	var details []ruckSacks
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var sackDetail ruckSacks
		line := scanner.Text()
		lineLength := len(line)

		if lineLength == 0 || lineLength%2 != 0 {
			fmt.Println("The input string cannot be divided equally.")
			return nil, errors.New("The input string cannot be divided ")
		}

		sackDetail.firstCompartment = line[:lineLength/2]
		sackDetail.secondCompartment = line[lineLength/2:]

		details = append(details, sackDetail)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return details, nil
}
