package main

import (
	"bufio"
	"fmt"
	"os"
)

func runSlope(lines []string, dx, dy int) int {
	xpos := 0
	ypos := 0
	n := 0
	for y, line := range lines {
		if ypos != y {
			continue
		}
		length := len(line)
		xposMod := xpos % length
		if line[xposMod] == '#' {
			n++
		}
		xpos += dx
		ypos += dy
	}
	return n
}

func main() {
	lines := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	n1 := runSlope(lines, 3, 1)
	fmt.Println("Part 1:", n1)

	n2 := runSlope(lines, 1, 1)
	n3 := runSlope(lines, 5, 1)
	n4 := runSlope(lines, 7, 1)
	n5 := runSlope(lines, 1, 2)
	fmt.Println("Part 2:", n1*n2*n3*n4*n5)
}
