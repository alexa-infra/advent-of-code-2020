package main

import (
	"fmt"
	"strconv"
)

func main() {
	const input = "523764819"
	ids := []int{}
	for _, ch := range input {
		n, _ := strconv.Atoi(string(ch))
		ids = append(ids, n)
	}
	state := map[int]int{}
	for i := 1; i < len(ids); i++ {
		state[ids[i-1]] = ids[i]
	}
	state[ids[len(ids)-1]] = ids[0]
	current := ids[0]

	max := 0
	for k := range state {
		if k > max {
			max = k
		}
	}

	getNext := func(a, b, c int) int {
		for i := current - 1; i > 0; i-- {
			if i != a && i != b && i != c {
				return i
			}
		}
		for i := max; i > 0; i-- {
			if i != a && i != b && i != c {
				return i
			}
		}
		panic("NO NEXT")
	}

	move := func() {
		// pick up
		a := state[current]
		b := state[a]
		c := state[b]
		state[current] = state[c]

		next := getNext(a, b, c)
		afterNext := state[next]
		state[next] = a
		state[a] = b
		state[b] = c
		state[c] = afterNext

		current = state[current]
	}

	for i := 0; i < 100; i++ {
		move()
	}
	p := state[1]
	text := ""
	for p != 1 {
		text += fmt.Sprint(p)
		p = state[p]
	}
	fmt.Println("Part 1:", text)

	for i := 10; i <= 1000*1000; i++ {
		ids = append(ids, i)
	}
	state = map[int]int{}
	for i := 1; i < len(ids); i++ {
		state[ids[i-1]] = ids[i]
	}
	state[ids[len(ids)-1]] = ids[0]
	current = ids[0]
	max = 1000 * 1000

	for i := 0; i < 10*1000*1000; i++ {
		move()
	}
	a := state[1]
	b := state[a]
	n2 := a * b
	fmt.Println("Part 2:", n2)
}
