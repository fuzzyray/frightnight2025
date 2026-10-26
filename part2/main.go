package main

/*
https://hyperskill.org/learn/daily/53136

Remember several properties of degrees:

360 degrees equal a full turn. So, for example, 390 degrees is the same as 30 degrees.

Counterclockwise (negative degrees) movement also falls somewhere on the 360 degrees scale. So, for example,
-40 degrees is the same as 320 degrees.

Given that information, your answer should be a number between 0 and 359.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part2")
	data, err := os.ReadFile("hyperskill-dataset-117295212.txt")
	if err != nil {
		log.Fatalln(err)
	}

	degreeEntries := strings.Split(strings.TrimRight(string(data), "\r\n"), ",")
	measurement := 0
	for _, degreeEntry := range degreeEntries {
		value, err := strconv.Atoi(degreeEntry)
		if err != nil {
			log.Fatalln(err)
		}
		measurement += value
	}
	measurement = measurement % 360
	if measurement < 0 {
		measurement = 360 + measurement
	}

	fmt.Println(measurement)
}
