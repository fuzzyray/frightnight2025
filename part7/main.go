package main

/*
https://hyperskill.org/learn/daily/53141

You will receive a decimal phone number as your dataset. Respond with a number of overflows and a 32-bit binary
unsigned integer in that order, divided by a single comma. If the decimal is too big to fit in the 32-bit binary,
increase the number of overflows and subtract decimal value at which overflow occurs from the phone number. Then try
to convert it again, until it does not overflow. If the binary has any leading zeroes, remove them.
*/

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcAnswer(number int) (int, int) {
	const maxValue = 1 << 32
	var overflow float64
	if number >= maxValue {
		overflow = math.Floor(float64(number) / float64(maxValue))
	}
	number = number - (maxValue * int(overflow))
	return int(overflow), number
}

func convertToBinary(number int) string {
	return fmt.Sprintf("%b", number)
}

func main() {
	fmt.Println("part7")
	data, err := os.ReadFile("hyperskill-dataset-117447412.txt")
	if err != nil {
		log.Fatalln(err)
	}

	phoneNumber, err := strconv.Atoi(strings.TrimRight(string(data), "\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
	overflow, number := calcAnswer(phoneNumber)
	fmt.Printf("%d,%b\n", overflow, number)
}
