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
	inc, curr, prev := 0, 0, 0

	for ifs.Scan() {
		curr, err = strconv.Atoi(ifs.Text())
		if err != nil {
			log.Fatal(err)
		}
		if curr > prev {
			inc++
		}
		prev = curr
	}
	fmt.Println("Total number of increases is ", inc-1)
}
