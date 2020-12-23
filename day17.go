package main

import (
	"fmt"
	"bufio"
	"os"
)

func iter3d(process func(x, y, z, w int)) {
	d := []int{ -1, 0, 1 }
	for _, x := range d {
		for _, y := range d {
			for _, z := range d {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				process(x, y, z, 0)
			}
		}
	}
}

func iter4d(process func(x, y, z, w int)) {
	d := []int{ -1, 0, 1 }
	for _, x := range d {
		for _, y := range d {
			for _, z := range d {
				for _, w := range d {
					if x == 0 && y == 0 && z == 0 && w == 0{
						continue
					}
					process(x, y, z, w)
				}
			}
		}
	}
}

func keyToCoord(key string) (int, int, int, int) {
	var x, y, z, w int
	_, err := fmt.Sscanf(key, "%d,%d,%d,%d", &x, &y, &z, &w)
	if err != nil {
		panic(err)
	}
	return x, y, z, w
}

func coordToKey(x, y, z, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
}

func run(data map[string]int, iter func(func (x, y, z, w int))) map[string]int {
	data2 := map[string]int{}
	for k, _ := range data {
		x, y, z, w := keyToCoord(k)
		numActive := 0
		iter(func(dx, dy, dz, dw int){
			x1, y1, z1, w1 := x + dx, y + dy, z + dz, w + dw
			key := coordToKey(x1, y1, z1, w1)
			_, ok := data[key]
			if ok {
				numActive++
			}
		})
		if numActive == 2 || numActive == 3 {
			// remains active
			data2[k] = 1
		}
	}
	for k, _ := range data {
		x, y, z, w := keyToCoord(k)
		iter(func(dx, dy, dz, dw int){
			x1, y1, z1, w1 := x + dx, y + dy, z + dz, w + dw
			key := coordToKey(x1, y1, z1, w1)
			_, ok := data[key]
			if !ok {
				numActive := 0
				iter(func(dx2, dy2, dz2, dw2 int){
					x2, y2, z2, w2 := x1 + dx2, y1 + dy2, z1 + dz2, w1 + dw2
					key := coordToKey(x2, y2, z2, w2)
					_, ok := data[key]
					if ok {
						numActive++
					}
				})
				if numActive == 3 {
					// becames active
					data2[key] = 1
				}
			}
		})
	}
	return data2
}

func main(){
	data := map[string]int{}
	j := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for i, ch := range line {
			if ch == '#' {
				key := coordToKey(i, j, 0, 0)
				data[key] = 1
			}
		}
		j++
	}
	orig_data := data

	for i := 0; i < 6; i++ {
		data = run(data, iter3d)
	}
	n1 := len(data)
	fmt.Println("Part 1:", n1)

	data = orig_data
	for i := 0; i < 6; i++ {
		data = run(data, iter4d)
	}
	n2 := len(data)
	fmt.Println("Part 2:", n2)
}
