package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	log.Printf("Count of safe reports are %d\n", solvePart1())
}

func solvePart1() (safeCount int) {
	list := readInput("inputs/day2.txt")
	for _, l := range list {
		//log.Printf("List %v is %t", l, isReportSafe(l))
		if isReportSafe(l) {
			//log.Printf("%v series is safe", l)
			safeCount++
		}
	}
	return
}

func isReportSafe(l []int) bool {
	//log.Printf("List is %v\n", l)
	isDecreasing := false
	if l[0] == l[1] {
		//log.Printf("%d:%d series is same...breaking", l[0], l[1])
		return false
	} else if l[0]-l[1] > 3 || l[1]-l[0] > 3 {
		//log.Printf("%d:%d series diff is greater than 3...breaking", l[0], l[1])
		return false
	} else if l[0] > l[1] {
		//log.Printf("%d:%d series is decreasing", l[0], l[1])
		isDecreasing = true
	} else if l[0] < l[1] {
		//log.Printf("%d:%d series is increasing", l[0], l[1])
	}
	for i := 1; i < len(l)-1; i++ {
		//log.Printf("Comparing %d and %d", l[i], l[i+1])
		if l[i] == l[i+1] {
			//log.Printf("%d:%d series is same..breaking", l[i], l[i+1])
			return false
		} else if l[i]-l[i+1] > 3 || l[i+1]-l[i] > 3 {
			//log.Printf("%d:%d series diff is greater than 3...breaking", l[i], l[i+1])
			return false
		} else if isDecreasing && (l[i] > l[i+1]) {
			//log.Printf("previous is decreasing and %d:%d series is decreasing..continuing...", l[i], l[i+1])
			//return false
		} else if !isDecreasing && (l[i] < l[i+1]) {
			//log.Printf("previous is increasing and %d:%d series is increasing..continuing", l[i], l[i+1])
		} else if isDecreasing && l[i] < l[i+1] {
			//log.Printf("previous is decreasing and %d:%d series is increasing..breaking", l[i], l[i+1])
			return false
		} else if !isDecreasing && l[i] > l[i+1] {
			//log.Printf("previous is increasing and %d:%d series is decreasing..breaking", l[i], l[i+1])
			return false
		}
	}
	return true
}
func readInput(fn string) (list [][]int) {
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
		d := strings.Split(ln, " ")
		l := make([]int, 0)
		for _, i := range d {
			n, err := strconv.Atoi(i)
			if err != nil {
				log.Fatalf("invalid number: %v", err)
			}
			l = append(l, n)
		}
		list = append(list, l)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
