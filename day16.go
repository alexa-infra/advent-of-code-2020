package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputRegex := regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	props := map[int][]*string{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			break
		}

		name := inputRegex.ReplaceAllString(text, "$1")
		a1, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))
		b1, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$3"))
		a2, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$4"))
		b2, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$5"))
		for i := a1; i <= b1; i++ {
			arr, ok := props[i]
			if !ok {
				props[i] = []*string { &name }
			} else {
				props[i] = append(arr, &name)
			}
		}
		for i := a2; i <= b2; i++ {
			arr, ok := props[i]
			if !ok {
				props[i] = []*string { &name }
			} else {
				props[i] = append(arr, &name)
			}
		}
	}

	scanner.Scan() // your ticket:
	scanner.Scan()
	text := scanner.Text()
	myTicketParts := strings.Split(text, ",")
	myTicket := []int{}
	for _, part := range myTicketParts {
		value, _ := strconv.Atoi(part)
		myTicket = append(myTicket, value)
	}

	scanner.Scan()
	scanner.Scan() // nearby tickets:

	tickets := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()
		ticketParts := strings.Split(text, ",")
		ticket := []int{}
		for _, part := range ticketParts {
			value, _ := strconv.Atoi(part)
			ticket = append(ticket, value)
		}
		tickets = append(tickets, ticket)
	}

	s1 := 0
	validTickets := [][]int{}
	for _, t := range tickets {
		valid := true
		for _, p := range t {
			_, ok := props[p]
			if !ok {
				s1 += p
				valid = false
			}
		}
		if !valid {
			continue
		}
		validTickets = append(validTickets, t)
	}
	//validTickets = append(validTickets, myTicket)

	tt := map[int][]string{}
	for i, _ := range myTicket {
		unique := map[string]int{}
		for _, t := range validTickets {
			value := t[i]
			match, _ := props[value]
			if len(unique) == 0 {
				for _, m := range match {
					unique[*m] = 1
				}
			} else {
				mm := map[string]int{}
				for _, m := range match {
					_, ok := unique[*m]
					if ok {
						mm[*m] = 1
					}
				}
				for k, _ := range unique {
					_, ok := mm[k]
					if !ok {
						delete(unique, k)
					}
				}
			}
		}
		arr := []string{}
		for k, _ := range unique {
			arr = append(arr, k)
		}
		tt[i] = arr
	}

	ttt := map[string]int{}
	for len(ttt) != 20 {
		for k, v := range tt {
			if len(v) == 1 {
				name := v[0]
				ttt[name] = k

				for k1, v1 := range tt {
					if k != k1 {
						arr := []string{}
						for _, z := range v1 {
							if z != name {
								arr = append(arr, z)
							}
						}
						tt[k1] = arr
					}
				}
				delete(tt, k)
			}
		}
	}
	s2 := 1
	for k, v := range ttt {
		if strings.HasPrefix(k, "departure") {
			s2 *= myTicket[v]
		}
	}
	fmt.Println("Part 1:", s1)
	fmt.Println("Part 2:", s2)
}
