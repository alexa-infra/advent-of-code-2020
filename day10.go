package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func pow(x, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func main() {
	numbers := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}

	sort.Sort(sort.IntSlice(numbers))

	diff := map[int]int{}
	currentLen := 0
	cuts := map[int]int{}
	for i, n := range numbers {
		d := n
		if i > 0 {
			d = n - numbers[i-1]
		}
		v, ok := diff[d]
		if ok {
			diff[d] = v + 1
		} else {
			diff[d] = 1
		}
		if d != 1 {
			if currentLen > 1 {
				vv, ok := cuts[currentLen]
				if ok {
					cuts[currentLen] = vv + 1
				} else {
					cuts[currentLen] = 1
				}
			}
			currentLen = 0
		} else {
			currentLen++
		}
	}
	if currentLen > 1 {
		vv, ok := cuts[currentLen]
		if ok {
			cuts[currentLen] = vv + 1
		} else {
			cuts[currentLen] = 1
		}
	}
	v1, _ := diff[1]
	v3, _ := diff[3]
	n1 := v1 * (v3 + 1)
	fmt.Println("Part 1:", n1)
	d2, _ := cuts[2]
	d3, _ := cuts[3]
	d4, _ := cuts[4]
	n2 := pow(7, d4) * pow(4, d3) * pow(2, d2)
	fmt.Println("Part 2:", n2)
}
