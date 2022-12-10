package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	value20  int
	value60  int
	value100 int
	value140 int
	value180 int
	value220 int
)

func main() {
	var (
		cycle        int
		xValue       = 1
		sum          int
		wantedCycles = []int{20, 60, 100, 140, 180, 220}
	)

	// readFile, err := os.Open("example.txt")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// if lo.Contains(wantedCycles, cycle) {
		// 	addToCycleValue(cycle, xValue)
		// 	fmt.Println("before all")
		// }
		if line == "noop" {
			fmt.Println("noop")
			cycle += 1
			if lo.Contains(wantedCycles, cycle) {
				// sum += cycle * xValue
				addToCycleValue(cycle, xValue)
				fmt.Println("inside noop")
			}
		} else {
			value, _ := strconv.Atoi(line[strings.LastIndex(line, " ")+1:])

			fmt.Println(value, cycle, xValue)
			cycle += 1
			if lo.Contains(wantedCycles, cycle) {
				addToCycleValue(cycle, xValue)
				// sum += cycle * xValue
				fmt.Println("inside first cycle of add")
			}
			cycle += 1
			if lo.Contains(wantedCycles, cycle) {
				addToCycleValue(cycle, xValue)
				// sum += cycle * xValue
				fmt.Println("inside last cycle of add")
			}
			xValue += value

		}

	}
	sum = value20 + value60 + value100 + value140 + value180 + value220
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Printing results:")

	fmt.Println("Value on cycle 20: ", value20)
	fmt.Println("Value on cycle 60: ", value60)
	fmt.Println("Value on cycle 100: ", value100)
	fmt.Println("Value on cycle 140: ", value140)
	fmt.Println("Value on cycle 180: ", value180)
	fmt.Println("Value on cycle 220: ", value220)
	fmt.Println("Total value: ", sum)

	fmt.Println("Total cycles: ", cycle)

}

func addToCycleValue(cycle int, xValue int) {
	if cycle == 20 {
		fmt.Println("Found cycle 20 ", xValue, cycle)
		value20 = xValue * cycle
	}
	if cycle == 60 {
		fmt.Println("Found cycle 60 ", xValue, cycle)
		value60 = xValue * cycle
	}
	if cycle == 100 {
		fmt.Println("Found cycle 100 ", xValue, cycle)
		value100 = xValue * cycle
	}
	if cycle == 140 {
		fmt.Println("Found cycle 140 ", xValue, cycle)
		value140 = xValue * cycle
	}
	if cycle == 180 {
		fmt.Println("Found cycle 180 ", xValue, cycle)
		value180 = xValue * cycle
	}
	if cycle == 220 {
		fmt.Println("Found cycle 220 ", xValue, cycle)
		value220 = xValue * cycle
	}
}
