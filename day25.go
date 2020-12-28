package main

import (
	"fmt"
)

func transform(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}
	return value
}

func getLoopSize(key, subjectNumber int) int {
	value := 1
	loopSize := 0
	for {
		loopSize++
		value = (value * subjectNumber) % 20201227
		if value == key {
			return loopSize
		}
	}
}

func main() {
	doorPublicKey := 9717666
	cardPublicKey := 20089533
	//doorLoopSize := getLoopSize(doorPublicKey, 7)
	cardLoopSize := getLoopSize(cardPublicKey, 7)
	encryptionKey := transform(doorPublicKey, cardLoopSize)
	fmt.Println("Part 1:", encryptionKey)
}
