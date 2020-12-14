package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"sort"
)

func findPairSumIn(arr []int, n int) int {
	for _, i := range(arr) {
		if i >= n {
			break
		}
		r := n - i
		for _, j := range(arr) {
			if j == r {
				return i * j
			} else if j > r {
				break
			}
		}
	}
	return 0
}

func main() {
        numbers := []int{}

        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
                n, _ := strconv.Atoi(scanner.Text())
                numbers = append(numbers, n)
        }

	sort.Sort(sort.IntSlice(numbers))
	n := findPairSumIn(numbers, 2020)
	fmt.Println("Part 1", n);

	for _, i := range numbers {
		if i >= 2020 {
			continue
		}
		r := 2020 - i
		n := findPairSumIn(numbers, r)
		if n != 0 {
			fmt.Println("Part 2", n * i);
			break
		}
	}
}
