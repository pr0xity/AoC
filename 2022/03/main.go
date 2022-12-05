package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var (
		stringOne      []string
		stringTwo      []string
		sum            int
		matchingLetter string
	)

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		stringOne = strings.Split(line[:len(line)/2], "")
		stringTwo = strings.Split(line[len(line)/2:], "")
		for _, x := range stringOne {
			for _, y := range stringTwo {
				if y == x {
					matchingLetter = x
					// fmt.Println(x)
					fmt.Println(sum)
					break
				}

			}
		}
		sum += strings.Index(priority, matchingLetter) + 1
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Printing results:")
	fmt.Println(sum)
	fmt.Println(strings.Index(priority, "a") + 1)

}
