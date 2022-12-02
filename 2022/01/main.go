package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

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
		elves []int
		elf   []int
	)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(fileScanner.Text())

		if len(line) < 1 {
			elves, elf = append(elves, lo.Sum(elf)), nil
			continue
		}
		elf = append(elf, lo.Must(strconv.Atoi(line)))
	}

	readFile.Close()

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	topThreeElves := lo.Sum(elves[0:3])
	fmt.Println("#1 elf:")
	fmt.Println(elves[0])
	fmt.Println("Top 3 elves:")
	fmt.Println(topThreeElves)
}
