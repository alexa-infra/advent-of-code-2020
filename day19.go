package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Rule struct {
	kind   int
	cargo0 string
	cargo1 string
	cargo2 string
	cargo3 string
	cargo4 string
}

func main() {
	re0 := regexp.MustCompile(`^(\d+): "(\w)"$`)
	re1 := regexp.MustCompile(`^(\d+): (\d+)$`)
	re2 := regexp.MustCompile(`^(\d+): (\d+) (\d+)$`)
	re3 := regexp.MustCompile(`^(\d+): (\d+) (\d+) (\d+)$`)
	re4 := regexp.MustCompile(`^(\d+): (\d+) (\d+) \| (\d+) (\d+)$`)
	re5 := regexp.MustCompile(`^(\d+): (\d+) \| (\d+)$`)

	rules := map[string]Rule{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else if re0.MatchString(line) {
			id := re0.ReplaceAllString(line, "$1")
			v0 := re0.ReplaceAllString(line, "$2")
			rule := Rule{0, v0, "", "", "", ""}
			rules[id] = rule
		} else if re1.MatchString(line) {
			id := re1.ReplaceAllString(line, "$1")
			v1 := re1.ReplaceAllString(line, "$2")
			rule := Rule{1, "", v1, "", "", ""}
			rules[id] = rule
		} else if re2.MatchString(line) {
			id := re2.ReplaceAllString(line, "$1")
			v1 := re2.ReplaceAllString(line, "$2")
			v2 := re2.ReplaceAllString(line, "$3")
			rule := Rule{2, "", v1, v2, "", ""}
			rules[id] = rule
		} else if re3.MatchString(line) {
			id := re3.ReplaceAllString(line, "$1")
			v1 := re3.ReplaceAllString(line, "$2")
			v2 := re3.ReplaceAllString(line, "$3")
			v3 := re3.ReplaceAllString(line, "$4")
			rule := Rule{3, "", v1, v2, v3, ""}
			rules[id] = rule
		} else if re4.MatchString(line) {
			id := re4.ReplaceAllString(line, "$1")
			v1 := re4.ReplaceAllString(line, "$2")
			v2 := re4.ReplaceAllString(line, "$3")
			v3 := re4.ReplaceAllString(line, "$4")
			v4 := re4.ReplaceAllString(line, "$5")
			rule := Rule{4, "", v1, v2, v3, v4}
			rules[id] = rule
		} else if re5.MatchString(line) {
			id := re5.ReplaceAllString(line, "$1")
			v1 := re5.ReplaceAllString(line, "$2")
			v2 := re5.ReplaceAllString(line, "$3")
			rule := Rule{5, "", v1, v2, "", ""}
			rules[id] = rule
		} else {
			panic("cant parse: " + line)
		}
	}
	var buildRule func(id string) string
	buildRule = func(id string) string {
		rule, ok := rules[id]
		if !ok {
			panic("rule not found " + id)
		}
		if rule.kind == 0 {
			return rule.cargo0
		}
		if rule.kind == 1 {
			return buildRule(rule.cargo1)
		}
		if rule.kind == 2 {
			return buildRule(rule.cargo1) + buildRule(rule.cargo2)
		}
		if rule.kind == 3 {
			return buildRule(rule.cargo1) + buildRule(rule.cargo2) + buildRule(rule.cargo3)
		}
		if rule.kind == 4 {
			part1 := buildRule(rule.cargo1) + buildRule(rule.cargo2)
			part2 := buildRule(rule.cargo3) + buildRule(rule.cargo4)
			return fmt.Sprintf("(%s|%s)", part1, part2)
		}
		if rule.kind == 5 {
			part1 := buildRule(rule.cargo1)
			part2 := buildRule(rule.cargo2)
			return fmt.Sprintf("(%s|%s)", part1, part2)
		}
		if rule.kind == 6 {
			part1 := buildRule(rule.cargo1)
			//part2 := buildRule(rule.cargo2)
			return fmt.Sprintf("(%s)+", part1)
		}
		if rule.kind == 7 {
			part1 := buildRule(rule.cargo1)
			part2 := buildRule(rule.cargo2)
			//part3 := buildRule(rule.cargo3)
			// x = ab|axb = ab|a(ab|a(ab|axb)b)b
			//              ab|aabb|aaabbb|...
			format := fmt.Sprintf("((%s){%%d}(%s){%%d})", part1, part2)
			p1 := fmt.Sprintf(format, 1, 1)
			p2 := fmt.Sprintf(format, 2, 2)
			p3 := fmt.Sprintf(format, 3, 3)
			p4 := fmt.Sprintf(format, 4, 4)
			p5 := fmt.Sprintf(format, 5, 5)
			return fmt.Sprintf("(%s|%s|%s|%s|%s)", p1, p2, p3, p4, p5)
		}
		panic("unknown kind")
	}
	regexpRule := buildRule("0")
	re := regexp.MustCompile("^" + regexpRule + "$")
	n1 := 0
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			n1++
		}
		lines = append(lines, line)
	}
	fmt.Println(n1)
	rules["8"] = Rule{6, "", "42", "", "", ""}
	rules["11"] = Rule{7, "", "42", "31", "", ""}

	regexpRule = buildRule("0")
	re = regexp.MustCompile("^" + regexpRule + "$")
	n2 := 0
	for _, line := range lines {
		if re.MatchString(line) {
			n2++
		}
	}
	fmt.Println(n2)
}
