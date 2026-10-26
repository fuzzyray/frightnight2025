package main

/*
https://hyperskill.org/learn/daily/53142

You will receive a dataset containing many words, written in the Old Gods language, split by a whitespace. You will
have to find one that does not have an anagram (meaning, you can't rearrange letters in the word to get another word
in the dataset). That will be your answer.
*/

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func main() {
	fmt.Println("part8")
	data, err := os.ReadFile("hyperskill-dataset-117477725.txt")
	if err != nil {
		log.Fatalln(err)
	}

	anagramWords := map[string][]string{}
	words := strings.Split(strings.TrimRight(string(data), "\r\n"), " ")

	for _, word := range words {
		anagram := sortString(word)
		anagramWords[anagram] = append(anagramWords[anagram], word)
	}

	for _, words := range anagramWords {
		if len(words) == 1 {
			fmt.Println(words[0])
		}
	}
}
