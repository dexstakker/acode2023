package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	m := map[string]int{}
	redCount := 12
	greenCount := 13
	blueCount := 14
	totalCount := 0
	// open file for reading
	// read line by line
	lines, err := readLines("data.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// print file contents
	for _, line := range lines {
		for k := range m {
			delete(m, k)
		}

		m["red"] = 0
		m["blue"] = 0
		m["green"] = 0
		gameId := 0
		printCode := 0

		// Break down each line as a Game Number followed
		res1 := strings.Split(line, ":")
		if printCode > 1 {
			//fmt.Println("Around colon = ", res1)
			os.Exit(1)
		}
		len1 := len(res1)
		if len1 >= 2 {
			idBreakdown := strings.Split(res1[0], " ")
			gameId, _ = strconv.Atoi(idBreakdown[1])
			//fmt.Println(res1[1])
		}
		foundViable := true
		res2 := strings.Split(res1[1], ";")
		//fmt.Println("res2 = ", res2)
		for _, res3 := range res2 {
			//fmt.Println("Dexline ", idx, res3)
			res4 := strings.Split(res3, ",")
			for _, choice := range res4 {
				t := strings.TrimSpace(choice)
				pieces := strings.Split(t, " ")
				color := pieces[1]
				count, _ := strconv.Atoi(pieces[0])
				m[color] = count
			}
			if m["red"] > redCount || m["blue"] > blueCount || m["green"] > greenCount {
				foundViable = false
				break
			}
		}
		if foundViable {
			//fmt.Println("Adding game id = ", gameId)
			totalCount += gameId
		}
	}
	fmt.Println("ANSWER = ", totalCount)
}
