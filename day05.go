package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func airplainPos(pos string) int {
	row := string(pos[:7])
	r := 0
	for i, v := range row {
		if string(v) == "B" {
			r = r | (1 << (6 - i))
		}
	}
	col := string(pos[7:])
	c := 0
	for i, v := range col {
		if string(v) == "R" {
			c = c | (1 << (2 - i))
		}
	}
	return r*8 + c
}

func main() {
	n1 := 0
	seats := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pos := scanner.Text()
		v := airplainPos(pos)
		seats = append(seats, v)
		if v > n1 {
			n1 = v
		}
	}
	fmt.Println("Part 1:", n1)
	sort.Sort(sort.IntSlice(seats))

	last := 0
	n2 := 0
	for _, i := range seats {
		if last != 0 {
			if last == i-2 {
				n2 = i - 1
			}
		}
		last = i
	}
	fmt.Println("Part 2:", n2)
}
