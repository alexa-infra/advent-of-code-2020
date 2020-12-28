package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//  __    __    __    __
// /  \__/  \__/  \__/  \
// \__/  \__/  \__/  \__/
// /  \__/  \__/  \__/  \
// \__/  \__/  \__/  \__/
// /  \__/  \__/  \__/  \
// \__/  \__/  \__/  \__/
// /  \__/  \__/  \__/  \
// \__/  \__/  \__/  \__/
// /  \__/  \__/  \__/  \
// \__/  \__/  \__/  \__/
//
// e, se, sw, w, nw, ne
//
// X: sw -> ne
// Y: nw -> se
// Z: w  -> e

type Pos struct {
	x, y, z int
}

func move(pos Pos, dir string) Pos {
	x, y, z := pos.x, pos.y, pos.z
	switch dir {
	case "e":
		return Pos{x - 1, y, z + 1}
	case "w":
		return Pos{x + 1, y, z - 1}
	case "ne":
		return Pos{x - 1, y + 1, z}
	case "sw":
		return Pos{x + 1, y - 1, z}
	case "se":
		return Pos{x, y - 1, z + 1}
	case "nw":
		return Pos{x, y + 1, z - 1}
	}
	panic("unknown direction")
}

func movePath(path []string, cells map[Pos]int) {
	pos := Pos{0, 0, 0}
	for _, dir := range path {
		pos = move(pos, dir)
	}
	v, ok := cells[pos]
	if !ok || v == 0 {
		cells[pos] = 1
	} else if v == 1 {
		cells[pos] = 0
	}
}

func getNeighbor(pos Pos, dir string, cells map[Pos]int) int {
	n := move(pos, dir)
	v, ok := cells[n]
	if !ok {
		return 0
	}
	return v
}

func allNeighbors(pos Pos) []Pos {
	directions := []string{"e", "w", "ne", "se", "nw", "sw"}
	neighbors := []Pos{}
	for _, dir := range directions {
		p := move(pos, dir)
		neighbors = append(neighbors, p)
	}
	return neighbors
}

func countBlackNeighbors(pos Pos, cells map[Pos]int) int {
	cc := 0
	for _, p := range allNeighbors(pos) {
		v, ok := cells[p]
		if ok && v == 1 {
			cc++
		}
	}
	return cc
}

func daily(cells map[Pos]int) map[Pos]int {
	newCells := map[Pos]int{}
	for pos := range cells {
		nBlack := countBlackNeighbors(pos, cells)
		if nBlack == 1 || nBlack == 2 {
			newCells[pos] = 1
		}
		for _, p := range allNeighbors(pos) {
			_, ok := cells[p]
			if !ok {
				nBlack := countBlackNeighbors(p, cells)
				if nBlack == 2 {
					newCells[p] = 1
				}
			}
		}
	}
	return newCells
}

func main() {
	inputRegex := regexp.MustCompile(`(e|se|sw|w|nw|ne)`)
	scanner := bufio.NewScanner(os.Stdin)
	cells := map[Pos]int{}
	for scanner.Scan() {
		text := scanner.Text()
		path := inputRegex.FindAllString(text, -1)
		movePath(path, cells)
	}
	for k, v := range cells {
		if v == 0 {
			delete(cells, k)
		}
	}
	n1 := len(cells)
	fmt.Println("Part 1:", n1)

	for i := 0; i < 100; i++ {
		cells = daily(cells)
	}
	n2 := len(cells)
	fmt.Println("Part 2:", n2)
}
