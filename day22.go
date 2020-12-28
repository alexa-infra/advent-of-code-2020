package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func game(deck1, deck2 []int) []int {
	player1 := append([]int{}, deck1...)
	player2 := append([]int{}, deck2...)

	for len(player1) > 0 && len(player2) > 0 {
		card1 := player1[0]
		player1 = player1[1:]
		card2 := player2[0]
		player2 = player2[1:]
		winner := card2 > card1
		if winner {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		} else {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		}
	}

	if len(player2) > 0 {
		return player2
	}
	return player1
}

func game2(depth int, deck1, deck2 []int) ([]int, bool) {
	records := map[string]int{}
	player1 := append([]int{}, deck1...)
	player2 := append([]int{}, deck2...)

	for len(player1) > 0 && len(player2) > 0 {
		decks := fmt.Sprint(player1, player2)
		if _, ok := records[decks]; ok {
			return player1, false
		}
		records[decks] = 1
		card1 := player1[0]
		player1 = player1[1:]
		card2 := player2[0]
		player2 = player2[1:]
		winner := card2 > card1
		if len(player1) >= card1 && len(player2) >= card2 {
			_, winner = game2(depth+1, player1[:card1], player2[:card2])
		}
		if winner {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		} else {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		}
	}

	if len(player2) > 0 {
		return player2, true
	}
	return player1, false
}

func main() {
	deck1 := []int{}
	deck2 := []int{}
	var currentDeck *[]int = nil
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Player 1:" {
			currentDeck = &deck1
		} else if text == "Player 2:" {
			currentDeck = &deck2
		} else if text == "" {
			continue
		} else {
			n, _ := strconv.Atoi(text)
			*currentDeck = append(*currentDeck, n)
		}
	}
	winner := game(deck1, deck2)
	n1 := 0
	for i, card := range winner {
		n1 += (len(winner) - i) * card
	}
	fmt.Println("Part 1:", n1)

	winner, _ = game2(0, deck1, deck2)
	n2 := 0
	for i, card := range winner {
		n2 += (len(winner) - i) * card
	}
	fmt.Println("Part 2:", n2)
}
