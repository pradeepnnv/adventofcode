package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	log.Printf("Total Difference is %d", solvePart1())
	log.Printf("Total Similarity score is %d", solvePart2())
}
func solvePart1() int {
	list1, list2 := readInput("inputs/day1.txt")

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

func solvePart2() (score int) {

	list1, list2 := readInput("inputs/day1.txt")
	//log.Println(list1, list2)
	sm := make(map[int]int)
	//prepare a score map of list2 items
	for _, v := range list2 {
		sm[v]++
	}
	//log.Println(sm)
	for i := 0; i < len(list1); i++ {
		s := sm[list1[i]]
		//log.Printf("similarity score for %d  is %d", list1[i], s*list1[i])
		score += s * list1[i]
	}
	return
}

func readInput(fn string) (list1, list2 []int) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	s := bufio.NewScanner(f)
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
	//return scanner
	return list1, list2
}
