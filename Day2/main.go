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
    usersMove := moves[1]
    var isUsersWin bool
    var isDraw bool

    if ((usersMove == "X" && opponentsMove == "A") || (usersMove == "Y" && opponentsMove == "B") || (usersMove == "Z" && opponentsMove == "C")) {
      isUsersWin = false
      isDraw = true
    } else if ((usersMove == "X" && opponentsMove == "C") || (usersMove == "Y" && opponentsMove == "A") || (usersMove == "Z" && opponentsMove == "B")) {
      isUsersWin = true
      isDraw = false
    } else {
      isUsersWin = false
      isDraw = false
    }

    if (isUsersWin) {
      score += 6
    } else if (isDraw) {
      score += 3
    }

    if (usersMove == "X") {
      score += 1
    } else if (usersMove == "Y") {
      score += 2
    } else if (usersMove == "Z") {
      score += 3
    }

  } else {
    return 0, errors.New("Invalid turnDetails")
  }

  return score, nil
}
