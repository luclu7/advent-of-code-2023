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
	inputTab := strings.Split(string(inputByte), "\n")

	// part1(inputTab)

	// os.Exit(0)

	results := 0

	// fmt.Println(inputTab, results)
	// lineLength := len(inputTab[0])

	re := regexp.MustCompile(`\*`)

	for line, inputLine := range inputTab {

		for _, v := range re.FindAllStringIndex(inputLine, -1) {
			fmt.Println("ligne", line+1, string(inputLine[v[0]]))

			linesToCheck := []int{}
			if line != 0 {
				linesToCheck = append(linesToCheck, line-1)
			}
			linesToCheck = append(linesToCheck, line)

			if line != len(inputTab)-1 {
				linesToCheck = append(linesToCheck, line+1)
			}

			// indexNumInCurrentLine := reNb.FindStringIndex(inputTab[line])
			// indexNumInCurrentLine := v
			// numberLine := inputTab[line]

			// lines
			concernedLines := []string{}
			for _, lineIndex := range linesToCheck {
				concernedLines = append(concernedLines, inputTab[lineIndex])
			}

			results += getNumbers(concernedLines, v[1])

			// affichage
			// for _, line := range concernedLines {
			// 	maxMoins := 0
			// 	maxPlus := 0
			// 	if indexNumInCurrentLine[0] > 2 {
			// 		maxMoins = 2
			// 	}
			// 	if indexNumInCurrentLine[1] < lineLength-2 {
			// 		maxPlus = 2
			// 	}

			// 	fmt.Print(line[indexNumInCurrentLine[0]-maxMoins : indexNumInCurrentLine[1]+maxPlus])
			// 	fmt.Println("", indexNumInCurrentLine)
			// }

			// fmt.Println("")

		}
	}

	fmt.Printf("Résultat total: %d\n", results)

	os.Exit(0)
	part1(inputTab)

}

func getNumbers(lines []string, gearPosition int) int {
	re := regexp.MustCompile(`\d{1,3}`)
	matches := []int{}

	for _, line := range lines {

		for _, v := range re.FindAllStringIndex(line, -1) {
			fmt.Print(v, " (", line[v[0]:v[1]], ") ", gearPosition, " - ")
			// chercher s'il y a un match qui est dans la même range que le gear
			fmt.Print("zone cherchée: \"", gearPosition, " à ", gearPosition+1, " donc ")
			isMatch := v[0] <= gearPosition && v[1] >= gearPosition-1
			if isMatch {
				fmt.Print("VRAI - ")
			} else {
				fmt.Print("FAUX - ")
			}

			if isMatch {
				fmt.Println("match !", line[v[0]:v[1]])
				nb, err := strconv.Atoi(line[v[0]:v[1]])
				if err != nil {
					panic(err.Error())
				}
				matches = append(matches, nb)
			} else {
				fmt.Println("")
			}

		}

	}

	total := 1
	if len(matches) > 1 {
		for _, v := range matches {
			total *= v
		}

		fmt.Printf("%d matchs: total de %d\n\n", len(matches), total)
		return total
	} else {
		fmt.Printf("%d matchs: pas de match\n\n", len(matches))
		return 0
	}

}

func part1(inputTab []string) {
	fmt.Println(inputTab)
	lineLength := len(inputTab[0])
	// input := string(inputByte)

	results := 0

	re := regexp.MustCompile("[0-9]{1,3}")
	for line, inputLine := range inputTab {

		for _, v := range re.FindAllStringIndex(inputLine, -1) {
			str := inputLine[v[0]:v[1]]

			fmt.Print(str)
			// compute the boundaries
			// boundaries est un tableau qui contient les abords de notre nombre
			boundaries := []string{}

			// reNb := regexp.MustCompile(str)
			// line := slices.IndexFunc(inputTab, func(c string) bool {
			// 	// cherche str (le nombre actuel) dans c (la ligne)
			// 	return reNb.FindString(c) == str
			// })

			fmt.Println(" l:", line+1)

			linesToCheck := []int{}
			if line != 0 {
				linesToCheck = append(linesToCheck, line-1)
			}
			linesToCheck = append(linesToCheck, line)

			if line != len(inputTab)-1 {
				linesToCheck = append(linesToCheck, line+1)
			}

			// indexNumInCurrentLine := reNb.FindStringIndex(inputTab[line])
			indexNumInCurrentLine := v
			// numberLine := inputTab[line]

			// append left and right character
			// if indexNumInCurrentLine[0] != 0 {
			// 	boundaries = append(boundaries, string(numberLine[int(indexNumInCurrentLine[0]-1)]))
			// }
			// if indexNumInCurrentLine[1] != lineLength {
			// 	boundaries = append(boundaries, string(numberLine[int(indexNumInCurrentLine[0]+1)]))
			// }

			for _, lineIndex := range linesToCheck {
				fmt.Print("Line ", lineIndex, ": ")
				line := inputTab[lineIndex]

				maxMoins := 0
				maxPlus := 0
				if indexNumInCurrentLine[0] > 0 {
					boundaries = append(boundaries, string(line[int(indexNumInCurrentLine[0]-1)]))
					maxMoins = 1
				}
				if indexNumInCurrentLine[1] < lineLength-1 {
					boundaries = append(boundaries, string(line[int(indexNumInCurrentLine[1])]))
					maxPlus = 1
				}

				fmt.Print(line[indexNumInCurrentLine[0]-maxMoins : indexNumInCurrentLine[1]+maxPlus])
				boundaries = append(boundaries, line[indexNumInCurrentLine[0]:indexNumInCurrentLine[1]])

				fmt.Println("", indexNumInCurrentLine)
			}

			// calculate every coordinates around the number
			// line by line

			boundariesStr := strings.Join(boundaries, "")

			fmt.Println(boundariesStr)

			reSymbol := regexp.MustCompile(`[^[0-9.]`)
			match := reSymbol.FindStringIndex(boundariesStr)

			if match != nil {
				fmt.Println("c'est bon !")
				nb, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				results += nb
			} else {
				fmt.Println("Pas bon !", str)
			}

			fmt.Println("")

		}
	}

	fmt.Println("Résultat: ", results)

}
