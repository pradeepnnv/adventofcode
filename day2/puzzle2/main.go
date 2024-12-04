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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ifs := bufio.NewScanner(f)

	// horizantal & depth
	h, d, a := 0, 0, 0

	for ifs.Scan() {
		input := strings.Split(ifs.Text(), " ")
		v, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}
		switch input[0] {

		case "down":
			a += v
		case "up":
			a -= v
		case "forward":
			h += v
			d += a * v
		default:
			log.Println("Invalid value for ", input[0])
		}
	}
	fmt.Println("Final position is :", h*d)
}
