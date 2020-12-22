package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputRegex := regexp.MustCompile(`^(\w)(\d+)$`)

	// counterclockwise
	dirX := []int{ 1, 0, -1,  0 }
	dirY := []int{ 0, 1,  0, -1 }

	currentDir := 0

	posX := 0
	posY := 0

	move := func(dirname string, units int){
		if dirname == "L" {
			times := (units / 90) % 4
			if currentDir + times > 3 {
				currentDir = currentDir + times - 4
			} else {
				currentDir = currentDir + times
			}
		} else if dirname == "R" {
			times := (units / 90) % 4
			if currentDir - times < 0 {
				currentDir = currentDir + 4 - times
			} else {
				currentDir = currentDir - times
			}
		} else {
			dir := -1
			if dirname == "F" {
				dir = currentDir
			} else if dirname == "E" {
				dir = 0
			} else if dirname == "N" {
				dir = 1
			} else if dirname == "W" {
				dir = 2
			} else if dirname == "S" {
				dir = 3
			}
			dx := dirX[dir]
			dy := dirY[dir]
			posX += dx * units
			posY += dy * units
		}
	}

	wpX := 10
	wpY := 1

	pX := 0
	pY := 0

	move2 := func(dirname string, units int){
		if dirname == "F" {
			pX += wpX * units
			pY += wpY * units
		} else if dirname == "L" {
			times := (units / 90) % 4
			for t := 0; t < times; t++ {
				wpX, wpY = -wpY, wpX
			}
		} else if dirname == "R" {
			times := (units / 90) % 4
			for t := 0; t < times; t++ {
				wpX, wpY = wpY, -wpX
			}
		} else {
			dir := -1
			if dirname == "E" {
				dir = 0
			} else if dirname == "N" {
				dir = 1
			} else if dirname == "W" {
				dir = 2
			} else if dirname == "S" {
				dir = 3
			}
			dx := dirX[dir]
			dy := dirY[dir]
			wpX += dx * units
			wpY += dy * units
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		ch := inputRegex.ReplaceAllString(text, "$1")
		n, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))

		move(ch, n)
		move2(ch, n)
	}
	m1 := abs(posX) + abs(posY)
	fmt.Println("Part 1:", m1)
	m2 := abs(pX) + abs(pY)
	fmt.Println("Part 2:", m2)
}

