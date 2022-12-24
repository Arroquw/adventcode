package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type supplies [][]string

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	day5(input)
}

func day5(input []string) {
	var c supplies
	var columns supplies
	var columns2 supplies
	var result string
	var result2 string
	for _, s := range input {
		if !strings.Contains(s, "move") {
			if s == "" {
				columns = transpose(c)
				columns2 = transpose(c)
			}
			if !strings.Contains(s, "[") {
				continue
			}
			//fmt.Println(temp)
			temp := replace(s)
			rowSplit := strings.Split(temp, "-")
			c = append(c, rowSplit)
						continue
		}
		temp := strings.Split(s, " ")
		count, _ := strconv.ParseInt(temp[1], 10, 32)
		source, _ := strconv.ParseInt(temp[3], 10 , 32)
		dest, _ := strconv.ParseInt(temp[5], 10 , 32)

		columns2.move(count, source, dest)
		columns.moveStack(count, source, dest)
	}
	for _, i := range columns {
		result += i[0]
	}
	for _, i := range columns2 {
		result2 += i[0]
	}
	temp2 := strings.ReplaceAll(result2, "[", "")
	fmt.Println(strings.ReplaceAll(temp2, "]", ""))
	temp := strings.ReplaceAll(result, "[", "")
	fmt.Println(strings.ReplaceAll(temp, "]", ""))
}

func (c *supplies) move(count int64, source int64, dest int64) {
	for i := int64(0); i < count; i++ {
		(*c)[dest-1] = append([]string{(*c)[source-1][0]}, (*c)[dest-1]...)
		(*c)[source-1] = (*c)[source-1][1:]
	}
}

func (c *supplies) moveStack(count int64, source int64, dest int64) {
	temp := make([]string, len((*c)[source-1][:count]))
	copy(temp, (*c)[source-1][:count])
	(*c)[dest-1] = append(temp, (*c)[dest-1]...)
	(*c)[source-1] = (*c)[source-1][count:]
}

func transpose(input supplies) supplies {
	result := make(supplies, len(input[0]))
	for _, row := range input {
		for j, value := range row {
			if !strings.Contains(value, "   ") {
				result[j] = append(result[j], value)
			}
		}
	}
	return result
}

func replace(input string) string {
	temp := []rune(input)
	count := 0
	for i, r := range temp {
		if r == ' ' {
			count++
		}
		if r == ']' || r == '[' {
			count = 3
		}
		if count%4 == 0 && r == ' ' {
			temp[i] = '-'
			count = 0
		}
	}
	return string(temp)
}

