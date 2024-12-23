package main

import (
	"bufio"
	"fmt"
	"os"

	"log"

	"github.com/pradeepnnv/adventofcode/2023/day2"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		fmt.Println("Input file needs to be provided as argument")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error while reading input %s", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}

	// fmt.Printf("Part1 output is %d\n", day1.CaliberationValuePart1(input))
	// fmt.Printf("Part2 output is %d\n", day1.CaliberationValuePart2(input))
	fmt.Printf("Part1 output is %d\n", day2.GameProcessorDay1(input))
	fmt.Printf("Part2 output is %d\n", day2.GameProcessorPart2(input))

}
