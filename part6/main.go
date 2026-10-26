package main

/*
You will receive a list of four-dimensional vectors, where each vector consists of four integers. Your goal is to
calculate the total distance traveled through these vectors in sequence:
1. Calculate the Euclidean distance between each consecutive pair of vectors (first to second, second to third, third
to fourth, etc.).
2. If any distance is not a whole number, round it up to the nearest integer.
3. Add all the rounded up distances together to get your final answer.

You can use the Euclidean distance formula for 4D vectors.
*/

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func convert(xs []string) []float64 {
	xf := make([]float64, len(xs))
	for i := range xs {
		if intVal, err := strconv.Atoi(xs[i]); err == nil {
			xf[i] = float64(intVal)
		} else {
			log.Fatalln(err)
		}
	}
	return xf
}

func calcEuclideanDistance(p1 []float64, p2 []float64) int {
	if len(p1) != len(p2) {
		log.Fatalln("length of p1 and p2 must be equal")
	}
	sumOfSquaredDifference := 0.0
	for i := range p1 {
		sumOfSquaredDifference += math.Pow(p1[i]-p2[i], 2)
	}
	distance := math.Sqrt(sumOfSquaredDifference)
	return int(math.Ceil(distance))
}

func calcTotalDistance(vectors [][]string) int {
	totalDistance := 0
	p1 := convert(vectors[0])
	for i := 1; i < len(vectors); i++ {
		p2 := convert(vectors[i])
		totalDistance += calcEuclideanDistance(p1, p2)
		p1 = p2
	}
	return totalDistance
}

func main() {
	fmt.Println("part6")
	file, err := os.Open("hyperskill-dataset-117415883.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	coords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	testCoords := [][]string{
		{"1", "1", "1", "1"},
		{"2", "2", "2", "2"},
		{"3", "3", "3", "3"},
		{"4", "4", "4", "4"},
	} // result = 6
	_ = testCoords

	fmt.Println(calcTotalDistance(coords))
}
