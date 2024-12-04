package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	calorieCount int
}

type ByCalorieCount []Elf

func (a ByCalorieCount) Len() int           { return len(a) }
func (a ByCalorieCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCalorieCount) Less(i, j int) bool { return a[i].calorieCount > a[j].calorieCount }

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(f)
	var elves []Elf
	var e Elf
	for inputScanner.Scan() {
		input := inputScanner.Text()
		if input == "" {
			elves = append(elves, e)
			e.calorieCount = 0
			continue
		}
		cal, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		e.calorieCount += cal
	}
	fmt.Println(elves)
	maxCalorie := 0
	for _, e := range elves {
		if e.calorieCount > maxCalorie {
			maxCalorie = e.calorieCount
		}
	}
	fmt.Println("Max Calorie count:", maxCalorie)

	sort.Sort(ByCalorieCount(elves))
	fmt.Println(elves)

	fmt.Println("Value of top three elves:", elves[0].calorieCount+elves[1].calorieCount+elves[2].calorieCount)
}
