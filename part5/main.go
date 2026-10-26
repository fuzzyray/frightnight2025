package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

/*
https://hyperskill.org/learn/daily/53139
You are doing well so far. Hang in there. This piano is special.
I retrieved it from some abandoned school. You see, it had some stupid bird-related puzzle attached to it.
I got rid of all that boring stuff, and reprogrammed it to use another sequence. Your clue is on the other side of
this note. Find the shortest sequence that contains all 7 notes, and play it to get the key.
*/

func main() {
	fmt.Println("part5")
	data, err := os.ReadFile("hyperskill-dataset-117376217.txt")
	if err != nil {
		log.Fatalln(err)
	}

	notes := strings.TrimRight(string(data), "\r\n")
	fmt.Println(getMinSequence(notes, "ABCDEFG"))
}

func getMinSequence(source string, target string) string {
	var sourceRunes = []rune(source)

	// frequency of needed characters
	var neededRunes = map[rune]int{}
	for _, c := range target {
		if _, ok := neededRunes[c]; !ok {
			neededRunes[c] = 1
		} else {
			neededRunes[c] += 1
		}
	}
	var need = len(neededRunes)

	var haveRunes = map[rune]int{}
	var left, have, start = 0, 0, 0
	var minLen = math.MaxInt

	for right, c := range sourceRunes {
		if _, ok := haveRunes[c]; !ok {
			haveRunes[c] = 1
		} else {
			haveRunes[c]++
		}
		if haveRunes[c] == neededRunes[c] {
			have++
		}

		for have == need {
			windowLen := right - left + 1
			if windowLen < minLen {
				minLen = windowLen
				start = left
			}

			leftRune := sourceRunes[left]
			haveRunes[leftRune]--
			if haveRunes[leftRune] < neededRunes[leftRune] {
				have--
			}
			left++
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return string(sourceRunes[start : start+minLen])
}
