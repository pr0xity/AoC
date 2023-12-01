package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

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
			n1_numbers []string
			n2_numbers []string
		)
		for i, char := range line {
			if unicode.IsDigit(char) {
				n1_numbers = append(n1_numbers, string(char))
				n2_numbers = append(n2_numbers, string(char))
			}
			for j, num := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], num) {
					n2_numbers = append(n2_numbers, strconv.Itoa(j+1))
				}
			}
		}
		n1 += lo.Must(strconv.Atoi(n1_numbers[0] + n1_numbers[len(n1_numbers)-1]))
		n2 += lo.Must(strconv.Atoi(n2_numbers[0] + n2_numbers[len(n2_numbers)-1]))
	}
	readFile.Close()

	fmt.Println(n1)
	fmt.Println(n2)
}
