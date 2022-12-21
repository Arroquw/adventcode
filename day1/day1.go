package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	day1(scanner)
}

func day1(scanner *bufio.Scanner) {
	var temp2 []int64
	var temp string
	max := (int64)(0)
	sum := make([]int64, 2)
	scanValid := true
	for {
		if scanValid {
			scanValid = scanner.Scan()
			temp = scanner.Text()
		}
		if temp != "" {
			num, _ := strconv.ParseInt(temp, 10, 64)
			temp2 = append(temp2, num)
			continue
		}
		for _, ijk := range temp2 {
			sum[0] += ijk
		}
		if sum[1] < sum[0] {
			max = sum[0]
			sum[1] = sum[0]
		} else {
			max = sum[1]
		}
		sum[0] = 0
		temp2 = make([]int64, 0)
		if !scanValid {
			break
		}
	}

	fmt.Println(max)
	if scanner.Err() != nil {
		fmt.Println("error")
	}

}
