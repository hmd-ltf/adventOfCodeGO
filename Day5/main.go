package main

import (
	"bufio"
	"os"
	"unicode"
)

type Move struct {
	moves     int
	fromCrate rune
	toCrate   rune
}
type CrateStack struct {
	label  rune
	crates []rune
}

func main() {
	crateStacks, moves := readFileData()
	print(crateStacks, moves)
}

func readFileData() ([]*CrateStack, []*Move) {

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var movesString []string
	var cratesString []string
	var crateLabels string

	readingCrates := true
	readingMoves := false

	for scanner.Scan() {
		data := scanner.Text()

		if data == "" {
			readingCrates = false
			readingMoves = true
			continue
		}

		if readingCrates {
			if isCrate(data[:3]) {
				cratesString = append(cratesString, data)
			} else {
				crateLabels = data
			}
		} else if readingMoves {
			movesString = append(movesString, data)
		}
	}

	crateStacks := mapDataToCrates(crateLabels, cratesString)
	moves := mapDataToMoves(movesString)

	return crateStacks, moves
}

func isCrate(data string) bool {
	return data[1] == ' ' || (data[0] == '[' && data[2] == ']' && unicode.IsLetter(rune(data[1])))
}

func mapDataToCrates(crateLabels string, cratesString []string) []*CrateStack {
	crateStacks := mapLabelsToCrateStacks(crateLabels)
	mapCratesToCrateStacks(cratesString, crateStacks)

	return crateStacks
}

func mapLabelsToCrateStacks(crateLabels string) []*CrateStack {
	crateStacks := make([]*CrateStack, 0)

	for _, crateLabel := range crateLabels {
		if crateLabel != ' ' {
			crate := &CrateStack{crateLabel, make([]rune, 0)}
			crateStacks = append(crateStacks, crate)
		}
	}

	return crateStacks
}

func mapCratesToCrateStacks(cratesString []string, crateStacks []*CrateStack) {
	for _, data := range cratesString {
		for i := 1; i < len(data); i += 4 {
			if data[i] != ' ' {
				stack := crateStacks[i/4]
				stack.crates = append(stack.crates, rune(data[i]))
			}
		}
	}
}

func mapDataToMoves(moveStrings []string) []*Move {
	moves := make([]*Move, 0)

	for _, moveString := range moveStrings {
		move := Move{
			moves:     int(moveString[5]),
			fromCrate: rune(moveString[12]),
			toCrate:   rune(moveString[17]),
		}

		moves = append(moves, &move)
	}

	return moves

}
