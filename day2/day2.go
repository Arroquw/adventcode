package main

import (
	"bufio"
	"os"
	"fmt"
)

var Moves = map[byte]map[byte]int {
	'A': map[byte]int { 'X': 4, 'Y': 8, 'Z': 3},
	'B': map[byte]int { 'X': 1, 'Y': 5, 'Z': 9},
	'C': map[byte]int { 'X': 7, 'Y': 2, 'Z': 6},
}

var Outcomes = map[byte]map[byte]int {
	'A': map[byte]int { 'X': 3, 'Y': 4, 'Z': 8},
	'B': map[byte]int { 'X': 1, 'Y': 5, 'Z': 9},
	'C': map[byte]int { 'X': 2, 'Y': 6, 'Z': 7},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	day2(scanner)
}

func day2(scanner *bufio.Scanner) {
	pt1score, pt2score := 0, 0
	for scanner.Scan() {
		temp := scanner.Text()
		pt1score += Moves[temp[0]][temp[2]]
		pt2score += Outcomes[temp[0]][temp[2]]
	}
	fmt.Println(pt1score, pt2score)
}


