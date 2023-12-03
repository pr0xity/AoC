package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var (
		n1 int
		n2 int
	)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var (
			ok         = true
			n2_numbers = make(map[string]int)
			power      = 1
		// n1_numbers []string
		)

		// split line into two parts, id and line, on :
		id := lo.Must(strconv.Atoi(strings.Split(strings.SplitN(line, ":", 2)[0], " ")[1]))
		line = strings.SplitN(line, ":", 2)[1]
		// split line into hands on ;
		hands := strings.Split(line, ";")
		// split hands into cubes on ,
		for _, hand := range hands {
			cubes := strings.Split(hand, ",")
			// split cubes into number and color on space
			for _, cube := range cubes {
				trimmedCube := strings.TrimSpace(cube)
				n, color := lo.Must(strconv.Atoi(strings.SplitN(trimmedCube, " ", 2)[0])), strings.SplitN(trimmedCube, " ", 2)[1]
				n2_numbers[color] = max(n2_numbers[color], n)
				// check if cubes contains more than 12 red, 13 green and 14 blue
				colorLimits := map[string]int{"red": 12, "green": 13, "blue": 14}
				if n > colorLimits[color] {
					ok = false
				}
			}
		}
		for _, v := range n2_numbers {
			power *= v
		}
		n2 += power
		// if not, add game id to n1
		if ok {
			n1 += id
		}

	}

	readFile.Close()

	fmt.Println(n1)
	fmt.Println(n2)
}
