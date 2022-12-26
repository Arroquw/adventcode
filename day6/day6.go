package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type window []rune

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		day6pt2(scanner.Text())
	}
}

func (s window) Push(v rune) window {
	temp := append(s, v)
	temp = temp[1:]
	return temp
}

func day6pt2(input string) {
	old := make(window, 14)
	duplicates := 0
	count := 0
	for _, c := range input {
		temp := string(old)
		for _, r := range old {
			if r != 0 {
				duplicates += strings.Count(temp, string(r))
			}
		}
		if duplicates == 14 && len(temp) >= 14{
			break
		}
		duplicates = 0
		count++
		old = old.Push(c)
	}
	fmt.Println(count)
}
