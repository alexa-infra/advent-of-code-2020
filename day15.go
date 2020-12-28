package main

import (
	"fmt"
)

func runGame(input []int, length int) int {
	data := map[int]int{}
	lastAge := 0
	add := func(pos, value int) {
		prev, ok := data[value]
		if ok {
			lastAge = pos - prev
		} else {
			lastAge = 0
		}
		data[value] = pos
	}
	lastAdded := 0
	for i := 0; i < length; i++ {
		if i < len(input) {
			v := input[i]
			add(i, v)
		} else {
			lastAdded = lastAge
			add(i, lastAge)
		}
	}
	return lastAdded
}

func main() {
	//fmt.Println(runGame([]int{0,3,6}, 2020))
	//fmt.Println(runGame([]int{1,3,2}, 2020))
	//fmt.Println(runGame([]int{2,1,3}, 2020))
	//fmt.Println(runGame([]int{1,2,3}, 2020))
	//fmt.Println(runGame([]int{2,3,1}, 2020))
	//fmt.Println(runGame([]int{3,2,1}, 2020))
	//fmt.Println(runGame([]int{3,1,2}, 2020))
	n1 := runGame([]int{13, 0, 10, 12, 1, 5, 8}, 2020)
	fmt.Println("Part 1:", n1)
	n2 := runGame([]int{13, 0, 10, 12, 1, 5, 8}, 30000000)
	fmt.Println("Part 2:", n2)
}
