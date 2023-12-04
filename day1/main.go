package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputByte, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err.Error())
	}

	input := strings.Split(string(inputByte), "\n")

	totalCalValue := 0

	for _, v := range input {
		numbers := []int{}
		// line := strings.Split(v, "")

		// for _, char := range line {
		// 	// est-ce ce que char est un nombre?
		// 	if _, err := strconv.Atoi(char); err == nil {
		// 		nb, err := strconv.Atoi(char)
		// 		if err != nil {
		// 			panic(err.Error())
		// 		}
		// 		numbers = append(numbers, nb)
		// 	}
		// }

		var re = regexp.MustCompile(`(?m)[0-9]|one|two|three|four|five|six|seven|eight|nine`)

		for i, match := range re.FindAllString(v, -1) {

			nb := parseNumber(match)
			if nb == -1 {
				fmt.Println("Attention! -1 détecté: ", v)
				break
			}

			fmt.Printf("Match %d: %s\n", i, match)

			numbers = append(numbers, nb)

		}

		if len(numbers) == 0 {
			fmt.Println("Pas de nombre trouvé:", v)
			continue
		}

		firstNum := numbers[0]
		lastNum := numbers[len(numbers)-1]
		calibrationVal := firstNum*10 + lastNum
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d: %d\n", numbers, calibrationVal)
		totalCalValue += calibrationVal
	}

	fmt.Printf("Total: %d\n", totalCalValue)
}

func parseNumber(nb string) int {

	switch nb {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		nb, err := strconv.Atoi(nb)
		if err != nil {
			panic(err.Error())
		}

		return nb

	}

}
