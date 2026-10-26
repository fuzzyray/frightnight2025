package main

/*
https://hyperskill.org/learn/daily/53144

There are 16 letters on each line. 8 left and 8 right. 14 of them, 7 left and 7 right, have an exact pair.
2 of them, on each side, do not. Your answer is the pairless letters.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	fmt.Println("part10")
	data := parseData("hyperskill-dataset-117529764.txt")
	for _, v := range data {
		left, right := v[0:8], v[8:]
		for _, letter := range left {
			if !strings.Contains(right, string(letter)) {
				fmt.Print(string(letter))
			}
		}
		for _, letter := range right {
			if !strings.Contains(left, string(letter)) {
				fmt.Print(string(letter))
			}
		}
	}
	fmt.Println()
}
