package main

import (
	"bufio"
	"os"
	"unicode"
)

type moves struct {
	moves     int
	fromCrate rune
	toCrate   rune
}
type crateStack struct {
	label  rune
	index  int
	crates []rune
}

func main() {
	crateStacks, moves := readFileData()
	print(crateStacks, moves)
}
func readFileData() ([]crateStack, []moves) {

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	crateStacks := make([]*crateStack, 0)

	for scanner.Scan() {
		readingCrates := true
		readingMoves := false
		data := scanner.Text()

		if data == "" {
			readingCrates = false
			readingMoves = true
		}

		if readingCrates {
			if isCrate(data[:3]) {
				for i := 1; i < len(data); i += 4 {
					if data[i] != ' ' {
						var stack crateStack

						if len(crateStacks) < i {
							stack.index = i
							stack.crates = make([]rune, 0)
						} else {
							stack = *crateStacks[i]
						}

						stack.crates = append(stack.crates, rune(data[i]))
						crateStacks = append(crateStacks, &stack)
					}
				}
			}

			print(crateStacks)
		} else if readingMoves {
		}
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return nil, nil
}

func isCrate(data string) bool {
	if data[0] == '[' && data[2] == ']' && unicode.IsLetter(rune(data[1])) {
		return true
	}

	if data[1] == ' ' {
		return true
	}

	return false
}
