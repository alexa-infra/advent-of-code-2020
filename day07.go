package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

type Bag struct {
	Name     string
	Cargo    []*Cargo
	InsideOf []*Bag
}

type Cargo struct {
	Bag   *Bag
	Count int
}

func main() {
	nodes := map[string]*Bag{}

	lines := []string{}
        scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)

		parts := strings.Split(text, " bags contain ")
		name := parts[0]
		nodes[name] = &Bag{ Name: name }
	}

	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		name := parts[0]
		node, _ := nodes[name]

		cargo := strings.TrimRight(parts[1], ".")
		if cargo != "no other bags" {
			cargo_parts := strings.Split(cargo, ", ")
			for _, cp := range cargo_parts {
				names := strings.Split(cp, " ")
				n, _ := strconv.Atoi(names[0])
				name1 := strings.Join(names[1:len(names)-1], " ")

				node1, ok := nodes[name1]
				if !ok {
					nodes[name1] = &Bag{ Name: name1 }
				}
				node.Cargo = append(node.Cargo, &Cargo{
					Count: n,
					Bag: node1,
				})
			}
		}
	}

	for _, node := range nodes {
		for _, another_node := range nodes {
			for _, cargo := range another_node.Cargo {
				if node == cargo.Bag {
					node.InsideOf = append(node.InsideOf, another_node)
				}
			}
		}
	}
	shiny_gold, _ := nodes["shiny gold"]
	unique := map[string]*Bag{ "shiny gold": shiny_gold }
	toprocess := []*Bag{ shiny_gold }
	for len(toprocess) > 0 {
		node := toprocess[0]
		toprocess = toprocess[1:]
		for _, pnode := range node.InsideOf {
			_, ok := unique[pnode.Name]
			if ok {
				continue
			}
			unique[pnode.Name] = pnode
			toprocess = append(toprocess, pnode)
		}
	}
	n1 := len(unique)-1
	fmt.Println("Part 1:", n1)

	n2 := 0
	next := []*Cargo { &Cargo{ Bag: shiny_gold, Count: 1} }
	for len(next) > 0 {
		node := next[0]
		next = next[1:]
		n2 += node.Count
		for _, cargo := range node.Bag.Cargo {
			next = append(next, &Cargo{ Bag: cargo.Bag, Count: cargo.Count * node.Count })
		}
	}
	n2 -= 1
	fmt.Println("Part 2:", n2)
}
