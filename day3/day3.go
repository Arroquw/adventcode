package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	day3pt1(input)
	day3pt2(input)
}

func day3pt1(input []string) {
	//	var chars []byte
	var shared []byte
	sum := 0
	found := false
	for _, s := range input {
		temp := s
		sz := (len(temp) / 2)
		for _, r := range temp[sz:] {
			if strings.Contains(temp[:sz], string(r)) && !found {
				//for index, s := range shared {
				//	if s == byte(r) {
				//		shared = append(shared[:index], shared[index+1:]...)
				//	}
				//}
				shared = append(shared, byte(r))
				found = true
			}
		}
		found = false
	}
	for _, c := range shared {
		sum += toPriority(c)
	}
	fmt.Println(sum)
}

func day3pt2(input []string) {
	var temp []string
	var badges []byte
	found := false
	sum := 0
	for index, _ := range input {
		if index > 1 {
			temp = []string{
				input[index],
				input[index-1],
				input[index-2]}
		}
		if (index+1)%3 == 0 && index > 0 {
			for _, t := range temp[2] {
				if strings.Contains(temp[0], string(t)) && strings.Contains(temp[1], string(t)) && !found {
					badges = append(badges, byte(t))
					found = true
				}
			}
			found = false
		}
	}
	for _, c := range badges {
		sum += toPriority(c)
	}
	fmt.Println(sum)
}

func toPriority(r byte) int {
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	} else {
		return int(r - 'a' + 1)
	}
}
