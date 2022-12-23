package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(day4(input))
}

func day4(input []string) (fullyOverlap int, overlap int) {
	fullyOverlap = 0
	overlap = 0
	for _, val := range input {
		pair := strings.Split(val, ",")
		ranges := append(strings.Split(pair[0], "-"), strings.Split(pair[1], "-")...)
		range0, _ := strconv.ParseInt(ranges[0], 10, 32)
		range1, _ := strconv.ParseInt(ranges[1], 10, 32)
		range2, _ := strconv.ParseInt(ranges[2], 10, 32)
		range3, _ := strconv.ParseInt(ranges[3], 10, 32)
		if (range0 <= range2 && range1 >= range3) || (range2 <= range0 && range3 >= range1) {
			fullyOverlap++
			overlap++
		} else if range1 >= range2 && range3 >= range0 {
			overlap++
		}
	}
	return fullyOverlap, overlap
}
