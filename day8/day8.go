package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
//	"strings"
)

func main() {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	day8(input)
}

func day8(input []string) {
	runes := make([][]rune, 0)
	for _, s := range input {
		r := make([]rune, 0)
		for _, b := range s {
			r = append(r, rune(b))
		}
		runes = append(runes, r)
	}
	totalCount := 0
	for index1, outer := range runes {
		for index2, inner := range outer {
			if (index1 > 0 && index1 < len(runes)-1) && (index2 > 0 && index2 < len(outer)-1) {
				count := iterate(outer[:index2], inner)
				count += iterate(outer[index2+1:], inner)
				fmt.Println("index1: ", index1, " index2: ", index2, " runes[index1]: " + string(runes[index1]))
				//convert columns to rows, iterate over converted
				fmt.Println(" -----------", count)
				totalCount += count
			} else {
				totalCount++
			}
		}
	}
	printRunes(runes)
	fmt.Println(totalCount)
}

func iterate(r []rune, i rune) int {
	fmt.Println("i: " + string(i))
	for _, z := range r {
		fmt.Print(string(z))
		if z >= i {
			fmt.Println()
			return 0
		}
	}
	fmt.Println()
	return 1
}

func printRunes(s [][]rune) {
	for _, outer := range s {
		for _, inner := range outer {
			fmt.Print(string(inner))
		}
		fmt.Println()
	}
}
