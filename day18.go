package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type expression []interface{}

func hasBrackets(expr expression) bool {
	for _, v := range expr {
		if vv, ok := v.(string); ok && vv == "(" {
			return true
		}
	}
	return false
}

func solveNoBrackets(expr expression, withPriority bool) int {
	if len(expr) == 1 {
		return expr[0].(int)
	} else if len(expr) == 3 {
		left, op, right := expr[0].(int), expr[1].(string), expr[2].(int)
		rv := 0
		if op == "+" {
			rv = left + right
		} else if op == "*" {
			rv = left * right
		} else {
			panic("not correct op")
		}
		return rv
	} else if len(expr) > 3 {
		if withPriority {
			for p := 1; p < len(expr); p += 2 {
				if expr[p].(string) == "+" {
					result := solveNoBrackets(expr[p-1:p+2], false)
					newExpr := expression{}
					newExpr = append(newExpr, expr[:p-1]...)
					newExpr = append(newExpr, result)
					newExpr = append(newExpr, expr[p+2:]...)
					return solveNoBrackets(newExpr, true)
				}
			}
		}
		result := solveNoBrackets(expr[0:3], false)
		newExpr := expression{result}
		newExpr = append(newExpr, expr[3:]...)
		return solveNoBrackets(newExpr, false)
	} else {
		panic("not correct expr")
	}
}

func resolveBrackets(expr expression, withPriority bool) expression {
	for i, v := range expr {
		if vv, ok := v.(string); ok && vv == "(" {
			n := 0
			for j := i + 1; j < len(expr); j++ {
				if vv, ok := expr[j].(string); ok && vv == "(" {
					n++
				}
				if vv, ok := expr[j].(string); ok && vv == ")" {
					if n > 0 {
						n--
					} else {
						inBrackets := expr[i+1 : j]
						value := solve(inBrackets, withPriority)
						newExpr := expression{}
						newExpr = append(newExpr, expr[:i]...)
						newExpr = append(newExpr, value)
						newExpr = append(newExpr, expr[j+1:]...)
						return newExpr
					}
				}
			}
			panic("no matching brackets 1")
		}
	}
	panic("no matching brackets 2")
}

func solve(expr expression, withPriority bool) int {
	for hasBrackets(expr) {
		expr = resolveBrackets(expr, withPriority)
	}
	return solveNoBrackets(expr, withPriority)
}

func convertExpr(parts []string) expression {
	expr := expression{}
	for _, part := range parts {
		intVal, err := strconv.Atoi(part)
		if err == nil {
			expr = append(expr, intVal)
		} else {
			expr = append(expr, part)
		}
	}
	return expr
}

func solveLine(line string, withPriority bool) int {
	re := regexp.MustCompile(`(\d+|[\+\-\*\(\)])`)
	parts := re.FindAllString(line, -1)
	expr := convertExpr(parts)
	return solve(expr, withPriority)
}

func main() {
	//fmt.Println(resolveBrackets(expression{ 1, "+", "(", 2, "+", 3, ")", "*", 9 }))
	//fmt.Println(solve(expression{ 1, "+", "(", 2, "+", 3, ")", "*", 9 }))
	//fmt.Println(solveNoBrackets(expression{ 1, "+", 1, "*", 8 }))
	//fmt.Println(solveLine("1 + 2 * 3 + 4 * 5 + 6"))
	//fmt.Println(solveLine("1 + (2 * 3) + (4 * (5 + 6))"))
	//fmt.Println(solveLine("2 * 3 + (4 * 5)")) // becomes 26.
	//fmt.Println(solveLine("5 + (8 * 3 + 9 + 3 * 4 * 3)")) // becomes 437.
	//fmt.Println(solveLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")) // becomes 12240.
	//fmt.Println(solveLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")) // becomes 13632
	//fmt.Println(solveLine("1 + 2 * 3 + 4 * 5 + 6", true))

	re := regexp.MustCompile(`(\d+|[\+\-\*\(\)])`)

	n1, n2 := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := re.FindAllString(line, -1)

		expr := convertExpr(parts)
		result := solve(expr, false)
		n1 += result

		result = solve(expr, true)
		n2 += result
	}
	fmt.Println("Part 1:", n1)
	fmt.Println("Part 2:", n2)
}
