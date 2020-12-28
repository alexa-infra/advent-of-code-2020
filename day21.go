package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func unique(list []string) []string {
	dict := map[string]int{}
	for _, item := range list {
		dict[item] = 1
	}
	rv := []string{}
	for k := range dict {
		rv = append(rv, k)
	}
	return rv
}

func intersect(a, b []string) []string {
	dict := map[string]int{}
	for _, item := range a {
		dict[item] = 1
	}
	rv := []string{}
	for _, item := range b {
		_, ok := dict[item]
		if ok {
			rv = append(rv, item)
		}
	}
	return rv
}

func difference(a, b []string) []string {
	dict := map[string]int{}
	for _, item := range a {
		dict[item] = 1
	}
	for _, item := range b {
		_, ok := dict[item]
		if ok {
			delete(dict, item)
		}
	}
	rv := []string{}
	for k := range dict {
		rv = append(rv, k)
	}
	return rv
}

func main() {
	inputRegex := regexp.MustCompile(`^((?:\w+ )+)\(contains (.*?)\)$`)

	allCandidates := map[string][]string{}
	allIngredients := map[string]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		parts := inputRegex.FindStringSubmatch(text)
		ingredients := strings.Split(strings.Trim(parts[1], " "), " ")
		allergens := strings.Split(strings.Trim(parts[2], " "), ", ")

		for _, allergen := range allergens {
			list, ok := allCandidates[allergen]
			if ok {
				allCandidates[allergen] = intersect(list, ingredients)
			} else {
				allCandidates[allergen] = ingredients
			}
		}
		for _, ingredient := range ingredients {
			count, ok := allIngredients[ingredient]
			if !ok {
				allIngredients[ingredient] = 1
			} else {
				allIngredients[ingredient] = count + 1
			}
		}
	}
	n1 := 0
	for ingredient, count := range allIngredients {
		found := false
		for _, cc := range allCandidates {
			for _, nn := range cc {
				if nn == ingredient {
					found = true
				}
			}
		}
		if !found {
			n1 += count
		}
	}
	fmt.Println("Part 1:", n1)

	confirmed := map[string]string{}
	for len(confirmed) < len(allCandidates) {
		confirmedIngredients := []string{}
		for _, ingredient := range confirmed {
			confirmedIngredients = append(confirmedIngredients, ingredient)
		}
		for allergen, ingredients := range allCandidates {
			if _, ok := confirmed[allergen]; !ok {
				set := difference(ingredients, confirmedIngredients)
				if len(set) == 1 {
					confirmed[allergen] = set[0]
				}
				allCandidates[allergen] = set
			}
		}
	}

	keys := []string{}
	for key := range confirmed {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	values := []string{}
	for _, key := range keys {
		values = append(values, confirmed[key])
	}
	n2 := strings.Join(values, ",")
	fmt.Println("Part 2:", n2)
}
