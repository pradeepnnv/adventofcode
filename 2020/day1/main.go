package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	nl := make([]int, 0)
	for s.Scan() {
		// fmt.Println(s.Text())
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
		}
		nl = append(nl, n)
	}
	fmt.Println(threeEntryProduct(nl))

}

func twoEntryProduct(nl []int) int {
	for i := 0; i < len(nl); i++ {
		for j := i + 1; j < len(nl); j++ {
			// fmt.Println("numbers are ", nl[i], nl[j])
			// fmt.Println("sum is :", nl[i]+nl[j])
			if nl[i]+nl[j] == 2020 {
				return nl[i] * nl[j]
			}
		}
	}
	return 0
}

func threeEntryProduct(nl []int) int {
	for i := 0; i < len(nl); i++ {
		for j := i + 1; j < len(nl); j++ {
			for k := j + 1; k < len(nl); k++ {
				fmt.Println("numbers are ", nl[i], nl[j])
				fmt.Println("sum is :", nl[i]+nl[j])
				if nl[i]+nl[j]+nl[k] == 2020 {
					return nl[i] * nl[j] * nl[k]
				}

			}
		}
	}
	return 0
}
