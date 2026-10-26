package main

/*
https://hyperskill.org/learn/daily/53138

A B C D
E F G H
I J K L
M N O P

DOWN,LEFT,RIGHT,UP,UP,RIGHT

You always start at top left corner (on letter A) for each input.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("part4")

	keypad := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
	}

	file, err := os.Open("hyperskill-dataset-117356296.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instructions := scanner.Text()
		fmt.Printf("%s", getButton(instructions, keypad))
	}
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getButton(instructions string, keypad [][]string) string {
	x, y := 0, 0
	xMax := len(keypad[0]) - 1
	yMax := len(keypad) - 1

	for _, instruction := range strings.Split(instructions, ",") {
		switch instruction {
		case "UP":
			if y == 0 {
				continue
			}
			y -= 1
		case "DOWN":
			if y == yMax {
				continue
			}
			y += 1
		case "LEFT":
			if x == 0 {
				continue
			}
			x -= 1
		case "RIGHT":
			if x == xMax {
				continue
			}
			x += 1
		default:
			continue
		}
	}
	return keypad[y][x]
}
