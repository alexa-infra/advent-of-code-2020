package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

type Map map[string]string

var (
	yearRegexp = regexp.MustCompile(`^(\d{4})$`)
	heightRegexp = regexp.MustCompile(`^(\d+)(cm|in)$`)
	colorRegexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eyeColorRegexp = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegexp = regexp.MustCompile(`^\d{9}$`)
)

func validateSimple(pass Map) bool {
	for k, v := range pass {
		if v == "" && k != "cid" {
			return false
		}
	}
	return true
}

func validateComplex(pass Map) bool {
	if pass["byr"] == "" || !yearRegexp.MatchString(pass["byr"]) {
		return false
	}
	byr, _ := strconv.Atoi(yearRegexp.ReplaceAllString(pass["byr"], "$1"))
	if byr < 1920 || byr > 2002 {
		return false
	}
	if pass["iyr"] == "" || !yearRegexp.MatchString(pass["iyr"]) {
		return false
	}
	iyr, _ := strconv.Atoi(yearRegexp.ReplaceAllString(pass["iyr"], "$1"))
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	if pass["eyr"] == "" || !yearRegexp.MatchString(pass["eyr"]) {
		return false
	}
	eyr, _ := strconv.Atoi(yearRegexp.ReplaceAllString(pass["eyr"], "$1"))
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	if pass["hgt"] == "" || !heightRegexp.MatchString(pass["hgt"]) {
		return false
	}
	hh, _ := strconv.Atoi(heightRegexp.ReplaceAllString(pass["hgt"], "$1"))
	mm := heightRegexp.ReplaceAllString(pass["hgt"], "$2")
	if mm == "cm" && (hh < 150 || hh > 193) {
		return false
	}
	if mm == "in" && (hh < 59 || hh > 76) {
		return false
	}
	if pass["hcl"] == "" || !colorRegexp.MatchString(pass["hcl"]) {
		return false
	}
	if pass["ecl"] == "" || !eyeColorRegexp.MatchString(pass["ecl"]) {
		return false
	}
	if pass["pid"] == "" || !pidRegexp.MatchString(pass["pid"]) {
		return false
	}
	return true
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)

	pass := Map{ "byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": "", "cid": "" }
	n1 := 0
	n2 := 0


        for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if validateSimple(pass) {
				n1 += 1
			}
			if validateComplex(pass) {
				n2 += 1
			}
			for k, _ := range pass {
				pass[k] = ""
			}
		} else {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				values := strings.Split(part, ":")
				name := values[0]
				pass[name] = values[1]
			}
		}
	}
	if validateSimple(pass) {
		n1 += 1
	}
	if validateComplex(pass) {
		n2 += 1
	}
	fmt.Println("Part 1:", n1)
	fmt.Println("Part 2:", n2)
}
