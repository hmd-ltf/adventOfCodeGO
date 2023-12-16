package main

import (
	"bufio"
	"os"
)

type moves struct {
	moves     int
	fromCrate rune
	toCrate   rune
}
type crateStack struct {
	label  rune
	crates []rune
}

func main() {
	crateStacks, moves := readFileData()
	print(crateStacks, moves)
}
func readFileData() ([]crateStack, []moves) {

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readingCrates := true
		readingMoves := false
		data := scanner.Text()

		if data == "" {
			readingCrates = false
			readingMoves = true
		}

		if readingCrates {

		} else if readingMoves {
		}
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return nil, nil
}
