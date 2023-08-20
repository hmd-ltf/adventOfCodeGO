package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
  calories := readFileData()
  caloriesPerPerson := sumAllCalories(calories)
  personCalories := maxCaloriesPerson(caloriesPerPerson)
  fmt.Println(personCalories)
}

func sumAllCalories(calories []string) []int {
  personalCaloriesCount := []int{0}
  personalCaloriesIndex := 0

  for _, value := range calories {
    if (len(value) != 0) {
      num, err := strconv.Atoi(value)

      if (err != nil) {
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
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    caloriesData = append(caloriesData, strings.Join(strings.Split(line, " "), ""))
  }

  return caloriesData
}

func maxCaloriesPerson(summedCalories []int) int {
  maxCalories := 0

  for _, value := range summedCalories {
    if(maxCalories < value) {
      maxCalories = value
    }
  }

  return maxCalories
}
