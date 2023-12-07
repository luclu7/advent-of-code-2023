package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputByte, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err.Error())
	}
	inputTab := strings.Split(string(inputByte), "\n")

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

	fmt.Println("Soit un total de ", total)
}
