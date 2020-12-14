package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	n1 := 0
	n2 := 0
	ans := map[string]int{}
	b := map[string]int{}
	i := 0
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			n1 += len(ans)
			n2 += len(b)
			ans = map[string]int{}
			b = map[string]int{}
			i = 0
		} else {
			a := map[string]int{}
			for _, p := range text {
				ch := string(p)
				ans[ch] = 1
				a[ch] = 1
				if i == 0 {
					b[ch] = 1
				}
			}
			for k := range b {
				_, ok := a[k]
				if !ok {
					delete(b, k)
				}
			}
			i += 1
		}
        }
	n1 += len(ans)
	n2 += len(b)
	fmt.Println("Part 1:", n1)
	fmt.Println("Part 2:", n2)
}
