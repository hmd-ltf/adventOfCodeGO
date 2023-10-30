package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type RuckSack struct {
	firstCompartment  string
	secondCompartment string
	commonItem        rune
	itemPriority      uint8
}

/* For part 1
	func calculateCommonItemPriority(sackDetails *RuckSack) {
		calculatedPrioritiesMap := make(map[rune]uint8)

MainLoop:

		for _, firstVal := range sackDetails.firstCompartment {
			for _, secondVal := range sackDetails.secondCompartment {
				if firstVal == secondVal {
					sackDetails.commonItem = firstVal
					sackDetails.itemPriority = calculatedPriority(firstVal, calculatedPrioritiesMap)
					break MainLoop
				}
			}
		}
	}
*/

func calculatedPriority(commonItem rune, priorityMap map[rune]uint8) uint8 {
	priority, isPriorityCalculated := priorityMap[commonItem]

	if !isPriorityCalculated {
		if commonItem <= 96 { // Its capital
			priority = uint8(commonItem) - 38
		} else {
			priority = uint8(commonItem) - 96
		}
		priorityMap[commonItem] = priority
	}

	return priority
}

func main() {
	var totalPriority uint16
	ruckSacksDetails := readFileData()
	calculatedPrioritiesMap := make(map[rune]uint8)

	/* for part 1
	   	for _, sackDetails := range ruckSacksDetails {
		calculateCommonItemPriority(&sackDetails)
		totalPriority = uint16(sackDetails.itemPriority) + totalPriority
	}
	*/

	for i := 0; i < len(ruckSacksDetails); i += 3 {
		commonItems := findCommonBadgeItem(ruckSacksDetails[i], ruckSacksDetails[i+1], ruckSacksDetails[i+2])
		badge := rune(commonItems[0])
		totalPriority = totalPriority + uint16(calculatedPriority(badge, calculatedPrioritiesMap))
	}

	print(totalPriority)
}

func findCommonBadgeItem(sack1 RuckSack, sack2 RuckSack, sack3 RuckSack) string {
	sack1Items := sack1.firstCompartment + sack1.secondCompartment
	sack2Items := sack2.firstCompartment + sack2.secondCompartment
	sack3Items := sack3.firstCompartment + sack3.secondCompartment

	return findCommonElements(findCommonElements(sack1Items, sack2Items), sack3Items)
}

func findCommonElements(str1 string, str2 string) string {
	existingCharacters := make(map[rune]bool)

	for _, char := range str1 {
		existingCharacters[char] = true
	}

	common := ""

	for _, char := range str2 {
		if existingCharacters[char] {
			common += string(char)
		}
	}

	return common
}

func readFileData() []RuckSack {
	var details []RuckSack
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var sackDetail RuckSack
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
