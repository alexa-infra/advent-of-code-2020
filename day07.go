package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	Name     string
	Cargo    []*cargoBag
	InsideOf []*bag
}

type cargoBag struct {
	Bag   *bag
	Count int
}

func main() {
	nodes := map[string]*bag{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, " bags contain ")
		name := parts[0]

		node, ok := nodes[name]
		if !ok {
			node = &bag{Name: name}
			nodes[name] = node
		}

		cargo := strings.TrimRight(parts[1], ".")
		if cargo != "no other bags" {
			cargoParts := strings.Split(cargo, ", ")
			for _, cargoPart := range cargoParts {
				names := strings.Split(cargoPart, " ")
				n, _ := strconv.Atoi(names[0])
				cargoName := strings.Join(names[1:len(names)-1], " ")

				cargoNode, ok := nodes[cargoName]
				if !ok {
					cargoNode = &bag{Name: cargoName}
					nodes[cargoName] = cargoNode
				}
				node.Cargo = append(node.Cargo, &cargoBag{
					Count: n,
					Bag:   cargoNode,
				})
			}
		}
	}

	for _, node := range nodes {
		for _, anotherNode := range nodes {
			for _, cargo := range anotherNode.Cargo {
				if node == cargo.Bag {
					node.InsideOf = append(node.InsideOf, anotherNode)
				}
			}
		}
	}
	shinyGold, _ := nodes["shiny gold"]
	unique := map[string]*bag{"shiny gold": shinyGold}
	toprocess := []*bag{shinyGold}
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
	n1 := len(unique) - 1
	fmt.Println("Part 1:", n1)

	n2 := 0
	next := []*cargoBag{&cargoBag{Bag: shinyGold, Count: 1}}
	for len(next) > 0 {
		node := next[0]
		next = next[1:]
		n2 += node.Count
		for _, cargo := range node.Bag.Cargo {
			next = append(next, &cargoBag{Bag: cargo.Bag, Count: cargo.Count * node.Count})
		}
	}
	n2--
	fmt.Println("Part 2:", n2)
}
