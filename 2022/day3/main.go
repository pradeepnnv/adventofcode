package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	prioritySum := 0
	inputScanner := bufio.NewScanner(f)
	var inputSlice []string
	for i := 0; i < 6; i++ {
		//read 3 items and construct the slice
		inputSlice = make([]string, 3)
		for j := 0; j < 3; j++ {
			inputScanner.Scan()
			inputSlice = append(inputSlice, inputScanner.Text())
			i++
		}
		prioritySum += priority(findCommonIteminSlice(inputSlice))
	}
	fmt.Printf("Total Priority for part 2 is %d\n", prioritySum)

}
func part1() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	prioritySum := 0
	inputScanner := bufio.NewScanner(f)
	for inputScanner.Scan() {
		input := inputScanner.Text()
		// fmt.Println(input)
		// fmt.Println(priority(findCommonItem(input)))
		// fmt.Println(priority('a'))
		// fmt.Println(priority('A'))
		prioritySum += priority(findCommonItem(input))
	}

	fmt.Printf("Total Priority for part 1 is %d\n", prioritySum)
}

// priority returns the priority of the item found
func priority(c rune) int {
	// fmt.Println("got ", c)
	if c >= 97 && c <= 122 {
		return int(c) - 97 + 1
	} else if c >= 65 && c <= 90 {
		return int(c) - 65 + 26 + 1
	}
	return int(c)
}

func findCommonItem(contents string) (item rune) {

	p1 := contents[:len(contents)/2]
	p2 := contents[len(contents)/2:]

	// fmt.Printf("Input is %s.\n Part 1 is %s. \n Part 2 is %s.\n", contents, p1, p2)

	for _, c := range p1 {
		if strings.ContainsRune(p2, c) {
			return c
		}
	}
	return
}

func findCommonIteminSlice(contents []string) (item rune) {

	fmt.Printf("Length of the slice is %d and contents of the slice are %v:", len(contents), contents)
	for _, c := range contents[0] {
		fmt.Printf("checking for the rune %c\n", c)
		if strings.ContainsRune(contents[1], c) && strings.ContainsRune(contents[2], c) {
			return c
		}
	}
	return
}
