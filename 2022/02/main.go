package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A = Rock		= X
// B = Paper	= Y
// C = Scissors	= Z

// Scores
// L = 0 = X
// D = 3 = Y
// W = 6 = Z

// Rock +1
// Paper +2
// Scissors +3

// Input text
// A X
// B Y
// C Z

func main() {

	readFile, err := os.Open("input.txt")
	var (
		totalPointsPartOne int
		totalPointsPartTwo int

		elf string
		you string
	)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		linearray := strings.Split(line, " ")
		elf = linearray[0]
		you = linearray[len(linearray)-1]

		fmt.Println(line)
		totalPointsPartOne = totalPointsPartOne + rpsPartOne(elf, you)
		totalPointsPartTwo = totalPointsPartTwo + rpsPartTwo(elf, you)
		// fmt.Println(elf)
		// fmt.Println(you)
		fmt.Println(totalPointsPartOne)
		fmt.Println(totalPointsPartTwo)
		linearray = nil
		// elf :=

	}

	fmt.Println("Solving parts")
	fmt.Println("--------------------------------")
	fmt.Println("Total points for part one:")
	fmt.Println(totalPointsPartOne)
	fmt.Println("Total points for part two:")
	fmt.Println(totalPointsPartTwo)

}

func rpsPartOne(elf string, you string) int {
	// Check for elf useing Rock
	if elf == "A" && you == "X" {
		fmt.Println("Draw")
		return 3 + 1
	}
	if elf == "A" && you == "Y" {
		fmt.Println("You win")
		return 6 + 2
	}
	if elf == "A" && you == "Z" {
		fmt.Println("Elf win")
		return 0 + 3
	}

	// Check for elf useing Paper
	if elf == "B" && you == "X" {
		fmt.Println("Elf win")
		return 0 + 1
	}
	if elf == "B" && you == "Y" {
		fmt.Println("Draw")
		return 3 + 2
	}
	if elf == "B" && you == "Z" {
		fmt.Println("You win")
		return 6 + 3
	}
	// Check for elf useing Scissors
	if elf == "C" && you == "X" {
		fmt.Println("You win")
		return 6 + 1
	}
	if elf == "C" && you == "Y" {
		fmt.Println("Elf win")
		return 0 + 2
	}
	if elf == "C" && you == "Z" {
		fmt.Println("Draw")
		return 3 + 3
	}
	return 0
}
func rpsPartTwo(elf string, you string) int {
	// Check for elf useing Rock
	if elf == "A" && you == "X" {
		fmt.Println("Elf win")
		return 0 + 3
	}
	if elf == "A" && you == "Y" {
		fmt.Println("Draw")
		return 3 + 1
	}
	if elf == "A" && you == "Z" {
		fmt.Println("You win")
		return 6 + 2
	}

	// Check for elf useing Paper
	if elf == "B" && you == "X" {
		fmt.Println("Elf win")
		return 0 + 1
	}
	if elf == "B" && you == "Y" {
		fmt.Println("Draw")
		return 3 + 2
	}
	if elf == "B" && you == "Z" {
		fmt.Println("You win")
		return 6 + 3
	}
	// Check for elf useing Scissors
	if elf == "C" && you == "X" {
		fmt.Println("Elf win")
		return 0 + 2
	}
	if elf == "C" && you == "Y" {
		fmt.Println("Draw")
		return 3 + 3
	}
	if elf == "C" && you == "Z" {
		fmt.Println("You win")
		return 6 + 1
	}
	return 0
}
