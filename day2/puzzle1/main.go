package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ifs := bufio.NewScanner(inputFile)

	// horizantal & depth
	h, d := 0, 0
	for ifs.Scan() {
		input := strings.Split(ifs.Text(), " ")
		v, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}
		switch input[0] {
		case "forward":
			h += v
		case "down":
			d += v
		case "up":
			d -= v
		default:
			log.Println("Invalid value for ", input[0])
		}
	}
	fmt.Println("Final position is :", h*d)
}
