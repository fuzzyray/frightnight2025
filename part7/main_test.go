package main

import (
	"testing"
)

func TestCalcAnswer(t *testing.T) {
	testNumber := 8076095357 // expected answer: 1,11100001010111110110111101111101 (3781128061)
	overflow, number := calcAnswer(testNumber)
	if overflow != 1 || number != 3781128061 {
		t.Errorf("Got: %d,%d Wanted: 1,3781128061\n", overflow, number)
	}
}

func TestBinary(t *testing.T) {
	testNumber := 8076095357
	_, number := calcAnswer(testNumber)
	wanted := "11100001010111110110111101111101"
	got := convertToBinary(number)
	if got != wanted {
		t.Errorf("Got: %s, Wanted: %s\n", got, wanted)
	}
}
