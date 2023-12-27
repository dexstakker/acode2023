package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// DEXNOTE: This the same code from Part A, just reading a different input file
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

	// open file for reading
	// read line by line
	lines, err := readLines("../data/final.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var times []int
	var distances []int
	for lineNum, line := range lines {
		// Grab the string staring at pos = 9
		if lineNum == 0 {
			res1 := line[8:]
			res1 = strings.TrimSpace(res1)
			res2 := strings.Fields(res1)

			for _, val := range res2 {
				val = strings.TrimSpace(val)
				intVal, _ := strconv.Atoi(val)
				times = append(times, intVal)
			}
		} else if lineNum == 1 {
			res1 := line[9:]
			res1 = strings.TrimSpace(res1)
			res2 := strings.Fields(res1)
			//fmt.Println(res2)
			for _, val := range res2 {
				intVal, _ := strconv.Atoi(val)
				distances = append(distances, intVal)
			}
		}
	}

	if len(times) != len(distances) {
		fmt.Println("Numbers don't sync")
		os.Exit(1)
	}

	var answers []int
	for i, _ := range times {
		t := 1
		count := 0
		for t < times[i] {
			curr := (times[i] - t) * t
			if curr > distances[i] {
				count += 1
				if i == 2 {
					fmt.Println("Where matching: ", t)
				}
			}
			t += 1
		}
		answers = append(answers, count)
	}
	total := 1
	for _, item := range answers {
		total *= item
	}
	fmt.Println(total)

}
