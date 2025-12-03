package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func SecretEntrancePart1(fName string, useSpecialMethod bool) int {
	c := 0
	d := 50
	input := readInput(fName)
	for _, s := range input {
		t, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}
		if t > 100 {
			if useSpecialMethod {
				c += t / 100
			}
			t = t % 100

		}
		switch s[0] {
		case 'R':
			if useSpecialMethod && d+t > 99 {
				c++
			}
			d += t
			if d > 99 {
				d -= 100
			}
		case 'L':
			if useSpecialMethod && d-t != 0 && d-t < 0 {
				c++
			}
			d -= t
			if d < 0 {
				d = 100 + d
			}
		default:
			log.Println("not turning dial at all")
		}
		if d == 0 && !useSpecialMethod {
			c++
		}
	}
	return c
}

func SecretEntrancePart2(fName string) int {
	return SecretEntrancePart1(fName, true)
}

func readInput(fName string) []string {
	f, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}
	return input
}
