package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"regexp"
	"strings"
)

func main() {
	inputRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)

	n1 := 0
	n2 := 0
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
		text := scanner.Text()
		min, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$1"))
		max, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))
		ch := inputRegex.ReplaceAllString(text, "$3")
		password := inputRegex.ReplaceAllString(text, "$4")

		// Part 1
		if count := strings.Count(password, ch); count >= min && count <= max {
			n1 += 1
		}

		// Part 2
		ch1 := string(password[min - 1])
		ch2 := string(password[max - 1])
		if ch1 == ch && ch2 != ch {
			n2 += 1
		} else if (ch1 != ch && ch2 == ch) {
			n2 += 1
		}
        }
	fmt.Println("Part 1:", n1);
	fmt.Println("Part 2:", n2);
}
