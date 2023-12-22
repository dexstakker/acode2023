package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func getPoints(p int) int {
	retval := 0
	if p == 0 {
		return 0
	} else {
		retval = (int)(math.Pow(2.0, (float64)(p-1)))
	}

	return retval
}

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	printCode := 1

	// open file for reading
	// read line by line
	lines, err := readLines("../data/data.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// print file contents
	matchCount := 0
	for _, line := range lines {
		//cardId := 0
		//Break the line on : to "Card nun" and numbers
		res1 := strings.Split(line, ":")
		if printCode > 1 {
			fmt.Println("Around colon = ", res1)
			//os.Exit(1)
		}

		//len1 := len(res1)
		//if len1 == 2 {
		//idBreakdown := strings.Split(res1[0], " ")
		//cardId, _ = strconv.Atoi(idBreakdown[1])
		//fmt.Println(res1[1])
		//}
		//fmt.Println("card_id = ", cardId, "data portion = ", res1[1])
		parts := strings.Split(res1[1], "|")
		matches := strings.TrimSpace(parts[0])
		array := strings.TrimSpace(parts[1])
		matchSlice := strings.Split(matches, " ")
		arraySlice := strings.Split(array, " ")
		// Thumb thru array Slice and see which elements in matchesSlice
		currCount := 0
		for _, val := range matchSlice {
			if len(val) == 0 {
				continue
			}
			if slices.Contains(arraySlice, val) {
				//fmt.Println("A = ", arraySlice, " CONTAINS val = ", val)
				currCount = currCount + 1
			}
		}
		points := getPoints(currCount)
		if currCount > 0 {
			matchCount = matchCount + points
			//fmt.Println("adding ", points)
		}
	}
	fmt.Println("Total count = ", matchCount)
}
