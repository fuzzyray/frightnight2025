package main

/*
https://hyperskill.org/learn/daily/53137
The clue is in unique letters
*/

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("part3")
	data, err := os.ReadFile("hyperskill-dataset-117333773.txt")
	if err != nil {
		log.Fatalln(err)
	}

	letters := strings.TrimRight(string(data), "\r\n")
	letterCounts := make(map[rune]int, 26)
	for _, letter := range letters {
		if _, ok := letterCounts[letter]; !ok {
			letterCounts[letter] = 1
		} else {
			letterCounts[letter] += 1
		}
	}
	for _, letter := range letters {
		if letterCounts[letter] == 1 {
			fmt.Printf("%s", string(letter))
		}
	}
	fmt.Println()
}
