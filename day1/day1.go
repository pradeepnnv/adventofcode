package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1() int {
	s := readInput("inputs/day1.txt")
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for s.Scan() {
		ln := strings.TrimSpace(s.Text())
		d := strings.Split(ln, "   ")
		if len(d) != 2 {
			log.Fatalf("wrong number of fields: %v in %s", len(d), ln)
		}

		n, err := strconv.Atoi(d[0])
		if err != nil {
			log.Fatalf("invalid number: %v", err)
		}
		list1 = append(list1, n)

		n, err = strconv.Atoi(d[1])
		if err != nil {
			log.Fatalf("invalid number: %v", err)
		}
		list2 = append(list2, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	//log.Println(list1, list2)
	sort.Ints(list1)
	sort.Ints(list2)
	//log.Println("Sorted lists", list1, list2)
	var diff int
	for i := 0; i < len(list1); i++ {
		var d int
		if list1[i] > list2[i] {
			d = list1[i] - list2[i]
		} else {
			d = list2[i] - list1[i]
		}
		//log.Printf("difference between %d and %d is %d\n", list1[i], list2[i], d)
		diff += d
	}
	return diff
}

func readInput(fn string) *bufio.Scanner {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner
}
