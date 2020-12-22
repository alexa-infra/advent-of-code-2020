package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	maskRegexp := regexp.MustCompile(`^mask = ([01X]{36})$`)
	opRegexp := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	maskAnd := int64(0)
	fixed := int64(0)
	maskOr := int64(0)
	floating := int64(0)
	data := map[int]int64{}
	data2 := map[int64]int64{}

	var writeFloating func(addr, mask, value int64)
	writeFloating = func(addr, mask, value int64){
		if mask == 0 {
			data2[addr] = value
			return
		}
		for i := 0; i < 36; i++ {
			bit := int64(1 << (35 - i))
			if bit & mask != 0 {
				writeFloating(addr | bit, mask & ^bit, value)
				writeFloating(addr & ^bit, mask & ^bit, value)
				return
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if maskRegexp.MatchString(text) {
			maskStr := maskRegexp.ReplaceAllString(text, "$1")
			maskAnd = 0
			fixed = 0
			maskOr = 0
			floating = 0
			for i, ch := range maskStr {
				bit := int64(1 << (35 - i))
				if ch == 'X' {
					maskAnd |= bit
					floating |= bit
				} else {
					if ch == '1' {
						fixed |= bit
						maskOr |= bit
					}
				}
			}
		} else if opRegexp.MatchString(text) {
			addr, _ := strconv.Atoi(opRegexp.ReplaceAllString(text, "$1"))
			value, _ := strconv.Atoi(opRegexp.ReplaceAllString(text, "$2"))
			data[addr] = int64(value) & maskAnd + fixed

			addr2 := int64(addr) | maskOr
			writeFloating(addr2, floating, int64(value))
		}
	}
	s1 := int64(0)
	for _, v := range data {
		s1 += v
	}
	fmt.Println("Part 1:", s1)

	s2 := int64(0)
	for _, v := range data2 {
		s2 += v
	}
	fmt.Println("Part 2:", s2)
}
