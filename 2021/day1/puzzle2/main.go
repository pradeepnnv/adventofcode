package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ifs := bufio.NewScanner(inputFile)

	input := make([]int, 2000)

	for i := 0; ifs.Scan(); i++ {
		input[i], _ = strconv.Atoi(ifs.Text())
	}

	inc, curr, prev := 0, 0, 0

	for i := 0; i < len(input)-2; i++ {
		curr = input[i] + input[i+1] + input[i+2]
		if curr > prev {
			inc++
		}
		prev = curr
	}
	fmt.Println("Total number of increases is ", inc-1)
}
