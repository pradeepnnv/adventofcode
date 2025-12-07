package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindInvalidIdSumPart1(input string) float64 {
	sum := 0.0
	ids := strings.Split(input, ",")
	for _, s := range ids {
		// fmt.Println(s)
		r := strings.Split(s, "-")
		// fmt.Printf("Start is %s and end %s\n", r[0], r[1])
		start, err := strconv.ParseFloat(r[0], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		end, err := strconv.ParseFloat(r[1], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i := start; i <= end; i++ {
			if !validateId(i) {
				// fmt.Printf("%0.0f is invalid id\n", i)
				sum += i
			}
		}
	}
	return sum
}

func validateId(i float64) bool {
	id := strconv.FormatFloat(i, 'f', 0, 64)
	// var pMap = make(map[string]int)

	// for j := 0; j < len(id)-1; j++ {
	// 	// fmt.Println("Checking for " + id[0:j+1] + " in " + id)
	// 	// if strings.Count(id, id[0:j+1]) == 2 {
	// 	// return false
	// 	pMap[id[0:j+1]] = strings.Count(id, id[0:j+1])
	// 	// }
	// }
	// fmt.Println(pMap)
	// for k, v := range pMap {
	// 	if len(k)%2 != 0 && v == 2 {
	// 		fmt.Println("found " + k + " in " + id + ". So declaring it as invalid")
	// 		return false
	// 	}
	// }

	// fmt.Println(id)

	if id[0:len(id)/2] == id[len(id)/2:] {
		// fmt.Printf("%s is an invalid id\n", id)
		return false
	}
	return true
}
