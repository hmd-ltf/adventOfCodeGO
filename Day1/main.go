package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calories := readFileData()
	caloriesPerPerson := sumAllCalories(calories)
	personCalories := findMaxOfTopThreeElfCalories(caloriesPerPerson)
	fmt.Println(personCalories)
}

func sumAllCalories(calories []string) []int {
	personalCaloriesCount := []int{0}
	personalCaloriesIndex := 0

	for _, value := range calories {
		if len(value) != 0 {
			num, err := strconv.Atoi(value)

			if err != nil {
				fmt.Println("Unable to parse string", err)
			} else {
				personalCaloriesCount[personalCaloriesIndex] = personalCaloriesCount[personalCaloriesIndex] + num
			}
		} else {
			personalCaloriesIndex++
			personalCaloriesCount = append(personalCaloriesCount, 0)
		}
	}

	return personalCaloriesCount
}

func readFileData() []string {
	var caloriesData []string
	file, _ := os.Open("data.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			print("unable to close the file due to: ", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		caloriesData = append(caloriesData, strings.Join(strings.Split(line, " "), ""))
	}

	return caloriesData
}

func findMaxOfTopThreeElfCalories(summedCalories []int) int {
	mostCaloriesByFirstElf := 0
	mostCaloriesBySecondElf := 0
	mostCaloriesByThirdElf := 0

	for _, value := range summedCalories {

		if mostCaloriesByFirstElf < value {
			mostCaloriesByThirdElf = mostCaloriesBySecondElf
			mostCaloriesBySecondElf = mostCaloriesByFirstElf
			mostCaloriesByFirstElf = value
		} else if mostCaloriesBySecondElf < value {
			mostCaloriesByThirdElf = mostCaloriesBySecondElf
			mostCaloriesBySecondElf = value
		} else if mostCaloriesByThirdElf < value {
			mostCaloriesByThirdElf = value
		}
	}

	return mostCaloriesByFirstElf + mostCaloriesBySecondElf + mostCaloriesByThirdElf
}
