package main

/*
https://hyperskill.org/learn/daily/53143

In this task, you will receive a map of the mansion as a dataset. # represents walls, . represents walkable floor,
P represents your starting point and G represents the garage. Your answer should be the shortest path in a format
like the following: UDLRRLDU, where U is moving Up, D is moving Down, L is moving Left and R is moving Right.

Some layouts may have multiple shortest paths, any of them should be accepted.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"part9/queue"
	"part9/stack"
)

type node struct {
	r        int
	c        int
	previous *node
}

func (n *node) ToString() string {
	return fmt.Sprintf("%v,%v", (*n).r, (*n).c)
}

func parseMap(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
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

func findStartAndEnd(mansion [][]rune) (*node, *node) {
	start, end := new(node), new(node)
	for row := range mansion {
		for col := range mansion[row] {
			if mansion[row][col] == 'P' {
				start.r = row
				start.c = col
			}
			if mansion[row][col] == 'G' {
				end.r = row
				end.c = col
			}
		}
	}
	return start, end
}

func getNextNodes(mansion [][]rune, n node) []node {
	row, col := n.r, n.c
	nextNodes := make([]node, 0)
	for _, vDir := range []int{-1, 1} {
		nextRow := row + vDir
		if nextRow < 0 || nextRow == len(mansion) {
			continue
		}
		if mansion[nextRow][col] != '#' {
			nextNodes = append(nextNodes, node{nextRow, col, nil})
		}
	}
	for _, hDir := range []int{-1, 1} {
		nextCol := col + hDir
		if nextCol < 0 || nextCol == len(mansion[0]) {
			continue
		}
		if mansion[row][nextCol] != '#' {
			nextNodes = append(nextNodes, node{row, nextCol, nil})
		}
	}
	return nextNodes
}

func bfs(mansion [][]rune, start node, end node) *node {
	var q queue.Queue
	discovered := make(map[string]bool)
	discovered[start.ToString()] = true
	q.Enqueue(start)

	for !q.IsEmpty() {
		v := q.Dequeue().(node)
		if v.ToString() == end.ToString() {
			return &v
		}
		for _, w := range getNextNodes(mansion, v) {
			if !discovered[w.ToString()] {
				discovered[w.ToString()] = true
				w.previous = &v
				q.Enqueue(w)
			}
		}
	}
	return nil
}

func getDirection(current node, next node) rune {
	switch current.r - next.r {
	case 1:
		return 'U'
	case -1:
		return 'D'
	}
	switch current.c - next.c {
	case 1:
		return 'L'
	case -1:
		return 'R'
	}

	return '.'
}

func getPath(start node, n node) string {
	s := stack.Stack{}
	next := n
	for next.previous != nil {
		s.Push(next)
		next = *next.previous
	}
	directions := make([]rune, 0)
	currentNode := start
	for !s.IsEmpty() {
		nextNode := s.Pop().(node)
		directions = append(directions, getDirection(currentNode, nextNode))
		currentNode = nextNode
	}

	return string(directions)
}

func main() {
	mansionMap := parseMap("hyperskill-dataset-117503623.txt")
	start, end := findStartAndEnd(mansionMap)
	answer := bfs(mansionMap, *start, *end)
	if answer != nil {
		path := getPath(*start, *answer)
		fmt.Println(path)
	}
}
