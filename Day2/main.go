package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func main()  {
  roundDetails := readFileData()
  totalScore := calculateScore(roundDetails)

  print(totalScore)
}

func calculateScore(roundsData []string) int {
  totalScore := 0

  for _, value := range roundsData {
    roundScore, err := findScore(value)

    if (err == nil) {
      totalScore += roundScore
    }
  }
  
  return totalScore
}

func readFileData() []string {
  var data []string
  file, _ := os.Open("data.txt")
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    data = append(data, line)
  }

  return data
}

func findScore(roundDetails string) (int, error) {
  moves := strings.Split(roundDetails, " ")
  score := 0

  if (len(moves) == 2) {
    opponentsMove := moves[0]
    roundResult := moves[1]
    var isUsersWin bool
    var isDraw bool

    if (roundResult == "X") {
      isUsersWin = false
      isDraw = false
    } else if (roundResult == "Y") {
      isUsersWin = false
      isDraw = true
      score += 3
    } else {
      isUsersWin = true
      isDraw = false
      score += 6
    }

    if (opponentsMove == "A") {
      if (isUsersWin) {
        score += 2
      } else if (isDraw) {
        score += 1
      } else {
        score += 3
      }
    } else if (opponentsMove == "B") {
      if (isUsersWin) {
        score += 3
      } else if (isDraw) {
        score += 2
      } else {
        score += 1
      }

    } else {
      if (isUsersWin) {
        score += 1
      } else if (isDraw) {
        score += 3
      } else {
        score += 2
      }

    }
  } else {
    return 0, errors.New("Invalid turnDetails")
  }

  return score, nil
}
