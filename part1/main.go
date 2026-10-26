package main

/*
https://hyperskill.org/learn/daily/53135

The code to the gates is four digits. It is the first, the second, the third and the fourth most common digit in
today's sequence, in that particular order. If there are equal number of appearances, higher digit should come first.
*/

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	fmt.Println("part1")
	const ASCII0 = 48
	const ASCII9 = 57

	data, err := os.ReadFile("hyperskill-dataset-117272632.txt")
	if err != nil {
		log.Fatalln(err)
	}

	digitCount := make([]int, 10)
	for _, char := range data {
		if char < ASCII0 || char > ASCII9 {
			continue
		}
		digitCount[char-ASCII0] += 1
	}

	sortedCount := make([]int, len(digitCount))
	copy(sortedCount, digitCount)
	slices.Sort(sortedCount)
	slices.Reverse(sortedCount)

	for i := 0; i < 4; i++ {
		for n := range digitCount {
			if sortedCount[i] == digitCount[n] {
				fmt.Printf("%d", n)
			}
		}
	}
	fmt.Println()
}
