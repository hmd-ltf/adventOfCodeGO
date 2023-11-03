package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RangeStruct struct {
	startingNumber int
	endingNumber   int
}

type RangePairs struct {
	firstPair  RangeStruct
	secondPair RangeStruct
}

func main() {
	rangePairs := readFileData()

	print(rangePairs)
}

func readFileData() []RangePairs {
	var rangePairs []RangePairs

	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rangePairs = append(rangePairs, prepareRangePairs(scanner.Text()))
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return rangePairs
}

func prepareRangePairs(data string) RangePairs {
	var rangePair RangePairs
	ranges := strings.Split(data, ",")

	rangePair.firstPair = prepareRangeStruct(ranges[0])
	rangePair.secondPair = prepareRangeStruct(ranges[1])

	return rangePair
}

func prepareRangeStruct(data string) RangeStruct {
	var rangeStruct RangeStruct

	ranges := strings.Split(data, "-")
	rangeStruct.startingNumber, _ = strconv.Atoi(ranges[0])
	rangeStruct.endingNumber, _ = strconv.Atoi(ranges[1])

	return rangeStruct
}
