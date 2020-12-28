package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"regexp"
)

const (
	U = 0
	R = 1
	D = 2
	L = 3
)

type Tile struct {
	id   int
	data [][]rune
}

type TilePos struct {
	x    int
	y    int
	rf   int
	*Tile
}

func (t *Tile) Rotate() {
	d := [][]rune{}
	for j := 0; j < 10; j++ {
		row := make([]rune, 10)
		for i := 0; i < 10; i++ {
			row[i] = t.data[i][9-j]
		}
		d = append(d, row)
	}
	t.data = d
}

func (t *Tile) Flip() {
	d := [][]rune{}
	for j := 0; j < 10; j++ {
		row := t.data[9-j]
		d = append(d, row)
	}
	t.data = d
}

func (t *Tile) Copy() *Tile {
	tile := &Tile{ id: t.id }
	for j := 0; j < 10; j++ {
		row := []rune{}
		for i := 0; i < 10; i++ {
			row = append(row, t.data[j][i])
		}
		tile.data = append(tile.data, row)
	}
	return tile
}

func (t *Tile) BorderRight() int {
	rv := 0
	for j := 0; j < 10; j++ {
		if t.data[j][9] == '#' {
			rv = rv | (1 << j)
		}
	}
	return rv
}

func (t *Tile) BorderLeft() int {
	rv := 0
	for j := 0; j < 10; j++ {
		if t.data[j][0] == '#' {
			rv = rv | (1 << (9 - j))
		}
	}
	return rv
}

func (t *Tile) BorderTop() int {
	rv := 0
	for i := 0; i < 10; i++ {
		if t.data[0][i] == '#' {
			rv = rv | (1 << i)
		}
	}
	return rv
}

func (t *Tile) BorderBottom() int {
	rv := 0
	for i := 0; i < 10; i++ {
		if t.data[9][i] == '#' {
			rv = rv | (1 << (9 - i))
		}
	}
	return rv
}

func (t *Tile) Borders() []int {
	return []int{ t.BorderTop(), t.BorderRight(), t.BorderBottom(), t.BorderLeft() }
}

func flip(v int) int {
	rv := 0
	for i := 0; i < 10; i++ {
		bit := 1 << i
		if v & bit == bit {
			rv = rv | (1 << (9 - i))
		}
	}
	return rv
}

func (tile *Tile) Orient() [][]int{
	borders := tile.Borders()
	t, r, b, l := borders[0], borders[1], borders[2], borders[3]
	t1, r1, b1, l1 := flip(b), flip(r), flip(t), flip(l)
	return [][]int{
		{ t, r, b, l },
		{ r, b, l, t },
		{ b, l, t, r },
		{ l, t, r, b },

		{ t1, r1, b1, l1 },
		{ r1, b1, l1, t1 },
		{ b1, l1, t1, r1 },
		{ l1, t1, r1, b1 },
	}
}

func (pos TilePos) AlignTile() *Tile {
	tile := pos.Tile.Copy()
	if pos.rf < 4 {
		for i := 0; i < pos.rf; i++ {
			tile.Rotate()
		}
	} else {
		tile.Flip()
		for i := 4; i < pos.rf; i++ {
			tile.Rotate()
		}
	}
	return tile
}

func solve(tiles []*Tile, initialOrient int) []TilePos {
	queue := []TilePos{}
	aligned := []TilePos{}
	notaligned := append([]*Tile{}, tiles...)

	findAlignedByPos := func (x, y int) (TilePos, bool) {
		for _, pos := range aligned {
			if pos.x == x && pos.y == y {
				return pos, true
			}
		}
		return TilePos{ 0, 0, 0, nil }, false
	}
	findAlignedBorder := func (x, y, d int) int {
		pos, found := findAlignedByPos(x, y)
		if !found {
			return -1
		}
		tile := pos.Tile
		borders := tile.Orient()[pos.rf]
		return flip(borders[d])
	}
	findNotAligned := func (x, y, t, r, d, l int) []TilePos {
		rv := []TilePos{}
		for _, tile := range notaligned {
			for i, bb := range tile.Orient() {
				if t != -1 && bb[U] != t {
					continue
				}
				if r != -1 && bb[R] != r {
					continue
				}
				if d != -1 && bb[D] != d {
					continue
				}
				if l != -1 && bb[L] != l {
					continue
				}
				rv = append(rv, TilePos{x, y, i, tile})
			}
		}
		return rv
	}
	dx := []int{ -1, 1,  0, 0 }
	dy := []int{  0, 0, -1, 1 }
	alignTile := func (pos TilePos) {
		aligned = append(aligned, pos)
		for i := 0; i < 4; i++ {
			x := pos.x + dx[i]
			y := pos.y + dy[i]
			_, found := findAlignedByPos(x, y)
			if !found {
				queue = append(queue, TilePos{x, y, 0, nil})
			}
		}
		for k, tile := range notaligned {
			if tile == pos.Tile {
				notaligned = append(notaligned[:k], notaligned[k+1:]...)
				break
			}
		}
	}

	alignTile(TilePos{ 0, 0, initialOrient, notaligned[8] })
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		t := findAlignedBorder(pos.x, pos.y-1, D)
		r := findAlignedBorder(pos.x+1, pos.y, L)
		d := findAlignedBorder(pos.x, pos.y+1, U)
		l := findAlignedBorder(pos.x-1, pos.y, R)

		matches := findNotAligned(pos.x, pos.y, t, r, d, l)
		if len(matches) == 0 {
		} else if len(matches) == 1 {
			alignTile(matches[0])
		} else {
			panic("MULTIPLE MATCHES FOUND")
		}
	}
	for i, pos := range aligned {
		aligned[i].Tile = pos.AlignTile()
	}
	return aligned
}

func findCorners(aligned []TilePos) (int, int, int, int) {
	minx, maxx, miny, maxy := 0, 0, 0, 0
	for i, pos := range aligned {
		if i == 0 {
			minx, maxx, miny, maxy = pos.x, pos.x, pos.y, pos.y
		} else {
			if pos.x < minx {
				minx = pos.x
			}
			if pos.x > maxx {
				maxx = pos.x
			}
			if pos.y < miny {
				miny = pos.y
			}
			if pos.y > maxy {
				maxy = pos.y
			}
		}
	}
	return minx, maxx, miny, maxy
}

func findDragons(aligned []TilePos) (int, int) {

	findAlignedByPos := func (x, y int) (TilePos, bool) {
		for _, pos := range aligned {
			if pos.x == x && pos.y == y {
				return pos, true
			}
		}
		return TilePos{ 0, 0, 0, nil }, false
	}

	minx, maxx, miny, maxy := findCorners(aligned)

	lines := []string{}
	for y := miny; y <= maxy; y++ {
		for j := 1; j < 9; j++ {
			text := ""
			for x := minx; x <= maxx; x++ {
				pos, _ := findAlignedByPos(x, y)
				tile := pos.Tile
				text += string(tile.data[j][1:9])
			}
			lines = append(lines, text)
		}
	}

	//for _, line := range lines {
	//	fmt.Println(line)
	//}
	//fmt.Println()

	draco0 := regexp.MustCompile(`..................#.`)
	draco1 := regexp.MustCompile(`#....##....##....###`)
	draco2 := regexp.MustCompile(`.#..#..#..#..#..#...`)

	nn := 0
	for _, line := range lines {
		for _, ch := range line {
			if ch == '#' {
				nn++
			}
		}
	}
	n2 := 0
	for j := 1; j < len(lines) - 1; j++ {
		line1 := lines[j]

		locations := draco1.FindAllStringIndex(line1, -1)
		if locations == nil {
			continue
		}
		for _, loc := range locations {
			a, b := loc[0], loc[1]
			line0 := lines[j-1][a:b]
			line2 := lines[j+1][a:b]
			if draco0.MatchString(line0) && draco2.MatchString(line2) {
				n2++
			}
		}
	}
	return nn - n2 * 15, n2
}

func main() {
	inputRegex := regexp.MustCompile(`^Tile (\d+):$`)

	tiles := []*Tile{}
	var currentTile *Tile = nil

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if inputRegex.MatchString(text) {
			id, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$1"))
			currentTile = &Tile{ id: id, }
			tiles = append(tiles, currentTile)
			continue
		}
		if currentTile != nil {
			row := []rune(text)
			currentTile.data = append(currentTile.data, row)
		}
	}

	aligned := solve(tiles, 0)

	findAlignedByPos := func (x, y int) (TilePos, bool) {
		for _, pos := range aligned {
			if pos.x == x && pos.y == y {
				return pos, true
			}
		}
		return TilePos{ 0, 0, 0, nil }, false
	}

	minx, maxx, miny, maxy := findCorners(aligned)

	corner1, _ := findAlignedByPos(minx, miny)
	corner2, _ := findAlignedByPos(maxx, miny)
	corner3, _ := findAlignedByPos(minx, maxy)
	corner4, _ := findAlignedByPos(maxx, maxy)
	n1 := corner1.Tile.id * corner2.Tile.id * corner3.Tile.id * corner4.Tile.id
	fmt.Println("Part 1:", n1)

	n2 := 0
	for i := 0; i < 8; i++ {
		aligned := solve(tiles, i)
		cc, dd := findDragons(aligned)
		if dd > 0 {
			n2 = cc
			break
		}
	}
	fmt.Println("Part 2:", n2)
}
