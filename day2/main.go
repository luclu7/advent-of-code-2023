package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// maxRed := 12
	// maxGreen := 13
	// maxBlue := 14

	inputByte, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err.Error())
	}

	input := strings.Split(string(inputByte), "\n")

	fmt.Println(input)

	total := 0

	for _, v := range input {
		inputLine := strings.Split(v, ": ")
		gameIDStr := strings.TrimPrefix(inputLine[0], "Game ")
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			panic(err.Error())
		}

		game := strings.Split(inputLine[1], "; ")

		fmt.Println(gameID)

		blue := 0
		red := 0
		green := 0

		for _, g := range game {
			a := strings.Split(g, ", ")
			for _, cubes := range a {
				// fmt.Println(cubes)
				b := strings.Split(cubes, " ")
				nbStr := b[0]
				nb, err := strconv.Atoi(nbStr)
				if err != nil {
					panic(err.Error())
				}
				color := strings.TrimSuffix(b[1], ",")

				// fmt.Printf("color: %s - %d \n", color, nb)

				switch color {
				case "red":
					if red < nb {
						red = nb
					}
				case "blue":
					if blue < nb {
						blue = nb
					}
				case "green":
					if green < nb {
						green = nb
					}

				}

			}
			// fmt.Println("")
		}

		// part 1
		// if blue > maxBlue || red > maxRed || green > maxGreen {
		// 	fmt.Println("Impossible !")
		// } else {
		// 	total += gameID
		// }

		power := green * blue * red

		fmt.Printf("power: %d, green %d blue %d red %d\n\n", power, green, blue, red)
		total += power
	}

	fmt.Printf("Total: %d\n", total)
}
