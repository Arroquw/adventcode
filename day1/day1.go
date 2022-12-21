package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var temp2 []int64
	max := (int64)(0)
	sum := make([]int64, 2)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := scanner.Text()
		if temp != "" {
			num, _ := strconv.ParseInt(temp, 10, 64)
			temp2 = append(temp2, num)
		} else {
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
		}
	}

	fmt.Println(max)
	if scanner.Err() != nil {
		fmt.Println("error")
	}
}
