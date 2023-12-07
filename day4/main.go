package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	matches int
	count   int
}

func main() {
	inputByte, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err.Error())
	}
	inputTab := strings.Split(string(inputByte), "\n")

	// part1(inputTab)

	// hashmap des cartes
	cards := map[int]Card{}

	// on remplit la hashmap cards
	for _, card := range inputTab {
		cardInfo := strings.Split(card, ": ")
		cardName := cardInfo[0]
		cardNumbers := strings.Split(cardInfo[1], " | ")
		winningNumbersStr := cardNumbers[0]
		ourNumbersStr := cardNumbers[1]

		winningNumbers := strings.Fields(strings.TrimSpace(winningNumbersStr))
		ourNumbers := strings.Fields(strings.TrimSpace(ourNumbersStr))
		ourNumbersArr := []int{}
		for _, v := range ourNumbers {
			nb, err := strconv.Atoi(v)
			if err != nil {
				panic(err.Error())
			}
			ourNumbersArr = append(ourNumbersArr, nb)
		}

		winningHashmap := map[int]bool{}
		for _, nbStr := range winningNumbers {
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				panic(err.Error())
			}
			winningHashmap[nb] = true
		}

		cardId, err := strconv.Atoi(strings.TrimSpace(strings.Split(cardName, "Card")[1]))
		if err != nil {
			panic(err.Error())
		}

		score := 0

		for _, nb := range ourNumbersArr {
			if winningHashmap[nb] {
				score++
			}
		}

		cards[cardId] = Card{
			cardId,
			score,
			1,
		}

	}

	for i := 1; i <= len(inputTab); i++ {
		card := cards[i]
		matches := card.matches
		for i := 0; i < card.count; i++ {

			for j := card.id + 1; j < card.id+matches+1; j++ {
				cardI := cards[j]
				cardI.count++
				cards[j] = cardI
			}
		}
	}

	score := 0
	for i := 1; i <= len(inputTab); i++ {
		score += cards[i].count
	}

	fmt.Println("score total: ", score)

}

func part1(inputTab []string) {
	total := 0

	for _, card := range inputTab {
		score := 0
		cardInfo := strings.Split(card, ": ")
		cardName := cardInfo[0]
		cardNumbers := strings.Split(cardInfo[1], " | ")
		winningNumbersStr := cardNumbers[0]
		ourNumbersStr := cardNumbers[1]

		winningNumbers := strings.Fields(strings.TrimSpace(winningNumbersStr))
		ourNumbers := strings.Fields(strings.TrimSpace(ourNumbersStr))

		// create hashmap
		winningHashmap := map[int]bool{}
		for _, nbStr := range winningNumbers {
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				panic(err.Error())
			}
			winningHashmap[nb] = true
		}

		for _, nbStr := range ourNumbers {
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				panic(err.Error())
			}
			if winningHashmap[nb] {
				fmt.Println("match !")
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		fmt.Println(cardName, ": ", score)
		total += score

	}

	fmt.Println("Soit un total de", total)
	os.Exit(0)
}
