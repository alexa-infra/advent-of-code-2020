package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkXMAS(list []int, value int) bool {
	min := -1
	max := -1
	for _, x := range list {
		if min == -1 || x < min {
			min = x
		}
		if max == -1 || x > max {
			max = x
		}
	}
	if value < min*2 || value > max*2 {
		return false
	}
	for i, x := range list {
		for j, y := range list {
			if i != j && x+y == value {
				return true
			}
		}
	}
	return false
}

func main() {
	var intList []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		intList = append(intList, value)
	}

	p := 25
	n1 := -1
	for i, x := range intList {
		if i >= p {
			if !checkXMAS(intList[i-p:i], x) {
				n1 = x
				break
			}
		}
	}
	fmt.Println("Part 1:", n1)

	for i, x := range intList {
		if x >= n1 {
			continue
		}
		s := x
		min := x
		max := x
		for j := i + 1; j < len(intList); j++ {
			y := intList[j]
			s += y
			if y < min {
				min = y
			}
			if y > max {
				max = y
			}
			if s == n1 {
				fmt.Println("Part 2:", min+max)
				return
			}
			if s > n1 {
				break
			}
		}
	}
}
