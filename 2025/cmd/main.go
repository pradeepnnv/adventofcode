package main

import (
	"log"

	"github.com/pradeepnnv/adventofcode/2025/day1"
)

func madin() {
	log.Printf("Password for Part1 is %d", day1.SecretEntrancePart1("../inputs/day1.txt", false))
	log.Printf("Password for Part2 is %d", day1.SecretEntrancePart2("../inputs/day1.txt"))
}
