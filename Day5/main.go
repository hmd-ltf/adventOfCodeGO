package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Move struct {
	moves     int
	fromCrate rune
	toCrate   rune
}

type CratesColumns struct {
	label  rune
	crates CrateStack
}

type CrateStack struct {
	items []rune
}

func (s *CrateStack) ReversePush(item rune) {
	s.items = append(s.items, item)
}
func (s *CrateStack) PushMultiple(item []rune) {
	s.items = append(s.items, item...)
}
func (s *CrateStack) Push(item rune) {
	s.items = append([]rune{item}, s.items...)
}
func (s *CrateStack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}
func (s *CrateStack) PopMultiple(number int) ([]rune, error) {
	if len(s.items) < number {
		return nil, fmt.Errorf("stack does not have enough elements")
	}

	top := s.items[:number]
	s.items = s.items[:len(s.items)-number]
	return top, nil
}
func (s *CrateStack) Peek() rune {
	if len(s.items) == 0 {
		return 0
	}

	return s.items[len(s.items)-1]
}

func main() {
	crateStacks, moves := readFileData()
	for _, move := range moves {
		//applyMoveOnCrates(crateStacks, move) // For 1
		applyMultipleMovesOnCrates(crateStacks, move) // For 2 still under work
	}

	print("Crates at top: ")
	for _, crateStack := range crateStacks {
		print(string(crateStack.crates.Peek()))
	}
}

func readFileData() ([]*CratesColumns, []*Move) {

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

func mapDataToCrates(crateLabels string, cratesString []string) []*CratesColumns {
	crateStacks := mapLabelsToCrateStacks(crateLabels)
	mapCratesToCrateStacks(cratesString, crateStacks)

	return crateStacks
}

func mapLabelsToCrateStacks(crateLabels string) []*CratesColumns {
	crateStacks := make([]*CratesColumns, 0)

	for _, crateLabel := range crateLabels {
		if crateLabel != ' ' {
			crate := &CratesColumns{crateLabel, CrateStack{}}
			crateStacks = append(crateStacks, crate)
		}
	}

	return crateStacks
}

func mapCratesToCrateStacks(cratesString []string, crateStacks []*CratesColumns) {
	for _, data := range cratesString {
		for i := 1; i < len(data); i += 4 {
			if data[i] != ' ' {
				stack := crateStacks[i/4]
				stack.crates.Push(rune(data[i]))
			}
		}
	}
}

func mapDataToMoves(moveStrings []string) []*Move {
	moves := make([]*Move, 0)
	pattern := `move (\d+) from (\d+) to (\d+)`
	re := regexp.MustCompile(pattern)

	for _, moveString := range moveStrings {
		matches := re.FindStringSubmatch(moveString)
		moveCount, _ := strconv.Atoi(matches[1])

		move := Move{
			moves:     moveCount,
			fromCrate: rune(matches[2][0]),
			toCrate:   rune(matches[3][0]),
		}

		moves = append(moves, &move)
	}

	return moves

}

func applyMoveOnCrates(crateStacks []*CratesColumns, move *Move) {
	fromCrate, toCrate, err := fetchCrateWithLabel(move.fromCrate, move.toCrate, crateStacks)

	if err == nil {
		for i := 0; i < move.moves; i++ {
			data, _ := fromCrate.crates.Pop()
			toCrate.crates.ReversePush(data)
		}
	} else {
		fmt.Println(err)
	}
}

func applyMultipleMovesOnCrates(crateStacks []*CratesColumns, move *Move) {
	fromCrate, toCrate, err := fetchCrateWithLabel(move.fromCrate, move.toCrate, crateStacks)

	if err == nil {
		if move.moves > 1 {
			data, _ := fromCrate.crates.PopMultiple(move.moves)
			toCrate.crates.PushMultiple(data)
		} else {
			data, _ := fromCrate.crates.Pop()
			toCrate.crates.ReversePush(data)
		}
	} else {
		fmt.Println(err)
	}
}

func fetchCrateWithLabel(fromCrateLabel rune, toCrateLabel rune, crateStack []*CratesColumns) (*CratesColumns, *CratesColumns, error) {
	var fromCrate *CratesColumns
	var toCrate *CratesColumns

	for _, crate := range crateStack {
		if crate.label == fromCrateLabel {
			fromCrate = crate
		} else if crate.label == toCrateLabel {
			toCrate = crate
		}
	}

	if fromCrate == nil || toCrate == nil {
		return nil, nil, errors.New("one or both crates not found")
	}

	return fromCrate, toCrate, nil
}
