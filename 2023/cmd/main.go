package main

import (
	"bufio"
	"fmt"
	"os"

	"log"

	"github.com/pradeepnnv/adventofcode/2023/day1"
)

func main() {
	f, err := os.Open("day1input.txt")
	if err != nil {
		log.Fatalf("error while reading input %s", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}

	fmt.Println(day1.CaliberationValuePart1(input))
}
