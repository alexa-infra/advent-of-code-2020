package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

type code struct {
	op    string
	value int
}

type program []*code;

func newCode(txt string) *code {
	parts := strings.Split(txt, " ")
	op := parts[0]
	value, _ := strconv.Atoi(parts[1])
	return &code{ op: op, value: value}
}

func (c *code) exec(pos int, acc int) (int, int) {
	switch c.op {
	case "acc":
		return pos + 1, acc + c.value
	case "jmp":
		return pos + c.value, acc
	case "nop":
		return pos + 1, acc
	}
	return pos + 1, acc
}

func (p program) run() (bool, int) {
	pos := 0
	acc := 0
	visited := map[int]int{}
	for {
		if pos >= len(p) {
			return true, acc
		}
		_, ok := visited[pos]
		if ok {
			return false, acc
		}
		visited[pos] = 1
		c := p[pos]
		pos, acc = c.exec(pos, acc)
	}
}

func main() {
	var prog program
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		prog = append(prog, newCode(line))
	}
	_, acc := prog.run()
	fmt.Println("Part 1:", acc);

	for i, c := range prog {
		var cc code
		if c.op == "acc" {
			continue
		}
		if c.op == "nop" {
			cc.op = "jmp"
			cc.value = c.value
		}
		if c.op == "jmp" {
			cc.op = "nop"
			cc.value = c.value
		}
		var p program
		if i > 0 {
			p = append(p, prog[:i]...)
		}
		p = append(p, &cc)
		if i + 1 < len(prog) {
			p = append(p, prog[i+1:]...)
		}
		exit, acc := p.run()
		if exit {
			fmt.Println("Part 2:", acc)
		}
	}
}
