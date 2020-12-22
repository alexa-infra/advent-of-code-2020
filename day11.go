package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDirX() []int {
	return []int{ 0, 1, 1,  1,  0, -1, -1, -1}
}

func getDirY() []int {
	return []int{ 1, 1, 0, -1, -1, -1,  0,  1}
}

func countAdj(data [][]rune, i, j, n int) int {

	dirX := getDirX()
	dirY := getDirY()

	yn := len(data)
	xn := len(data[0])
	rv := 0

	for k := 0; k < 8; k++ {
		x := dirX[k]
		y := dirY[k]

		i1 := i
		j1 := j

		for nn := 0; nn < n; nn++ {
			i1 += x
			j1 += y
			if i1 < 0 || i1 >= xn {
				break
			}
			if j1 < 0 || j1 >= yn {
				break
			}
			ch := data[j1][i1]
			if ch == '#' {
				rv++
				break
			} else if ch == 'L' {
				break
			}
		}
	}
	return rv
}

func step(data [][]rune, t, m int) ([][]rune, int) {
	rv := [][]rune{}
	n := 0
	for j, row := range data {
		nrow := []rune{}
		for i, ch := range row {
			if ch == '.' {
				nrow = append(nrow, '.')
			} else {
				c := countAdj(data, i, j, m)
				if ch == 'L' {
					if c == 0 {
						nrow = append(nrow, '#')
						n++
					} else {
						nrow = append(nrow, 'L')
					}
				} else if ch == '#' {
					if c >= t {
						nrow = append(nrow, 'L')
						n++
					} else {
						nrow = append(nrow, '#')
					}
				}
			}
		}
		rv = append(rv, nrow)
	}
	return rv, n
}

//func printmap(data [][]rune){
//	for _, row := range data {
//		fmt.Println(string(row))
//	}
//	fmt.Println()
//}

func main() {
	data := [][]rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		data = append(data, row)
	}
	orig_data := data

	n1 := 0
	for {
		ndata, n := step(data, 4, 1)
		if n == 0 {
			for _, row := range data {
				for _, ch := range row {
					if ch == '#' {
						n1++
					}
				}
			}
			break
		}
		data = ndata
	}
	fmt.Println("Part 1:", n1)

	data = orig_data
	n2 := 0
	for {
		ndata, n := step(data, 5, 100)
		if n == 0 {
			for _, row := range data {
				for _, ch := range row {
					if ch == '#' {
						n2++
					}
				}
			}
			break
		}
		data = ndata
	}
	fmt.Println("Part 2:", n2)
}
