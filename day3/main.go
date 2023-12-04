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

	part1(inputTab)

	os.Exit(0)

	results := 0

	fmt.Println(inputTab, results)
	lineLength := len(inputTab[0])

	re := regexp.MustCompile("\\*")
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
				if indexNumInCurrentLine[0] > 2 {
					boundaries = append(boundaries, string(line[int(indexNumInCurrentLine[0]-1)]))
					maxMoins = 2
				}
				if indexNumInCurrentLine[1] < lineLength-2 {
					boundaries = append(boundaries, string(line[int(indexNumInCurrentLine[1])]))
					maxPlus = 2
				}

				fmt.Print(line[indexNumInCurrentLine[0]-maxMoins : indexNumInCurrentLine[1]+maxPlus])
				boundaries = append(boundaries, line[indexNumInCurrentLine[0]:indexNumInCurrentLine[1]])

				fmt.Println("", indexNumInCurrentLine)
			}

			// calculate every coordinates around the number
			// line by line

			boundariesStr := strings.Join(boundaries, "")

			fmt.Println(boundariesStr)

			// find if there are two numbers

			fmt.Println("")

		}
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
