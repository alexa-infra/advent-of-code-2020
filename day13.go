package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	time, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	minWait := 0
	minBus := 0
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		busId, _ := strconv.Atoi(bus)
		wait := busId - (time % busId)
		if minBus == 0 || wait < minWait {
			minWait = wait
			minBus = busId
		}
	}
	n1 := minBus * minWait
	fmt.Println("Part 1:", n1)

	// each bus goes in (busId * i)
	// bus with offet should go in (t + offset) = (busN * i)
	// in the same time (as above), offset = busN - t % busN
	//
	// with two buses
	// t = bus_1 * i
	// offset = bus_2 - t % bus_2
	//
	// we can find i, which would give (bus_1 * i) % bus_2 = bus_2 - offset
	// by knowing i we also know t
	//
	// then adding third bus
	// offset = bus_3 - t % bus_3
	// to statisfy previous equations we need to add lcm(bus_1, bus_2)=bus_1*bus_2
	// because all buses are prime

	acc := int64(0)
	lcm := int64(1)
	for offset, bus := range buses {
		if bus == "x" {
			continue
		}
		bus32, _ := strconv.Atoi(bus)
		bus64 := int64(bus32)

		if offset != 0 {
			diff := (bus64 - int64(offset)) % bus64
			if diff < 0 {
				diff += bus64
			}
			for acc % bus64 != diff {
				acc += lcm
			}
		}
		lcm *= bus64
	}
	fmt.Println("Part 2:", acc)
}

