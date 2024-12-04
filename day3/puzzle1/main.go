package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ifs := bufio.NewScanner(f)
	p1Count,
		p2Count,
		p3Count,
		p4Count,
		p5Count, p6Count, p7Count, p8Count, p9Count, p10Count, p11Count, p12Count := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	for ifs.Scan() {
		for i := 0; i < 1000; i++ {
			t := ifs.Text()
			switch i {
			case 0:
				if t[i] == '1' {
					p1Count++
				}
			case 1:
				if t[i] == '1' {
					p2Count++
				}
			case 2:
				if t[i] == '1' {
					p3Count++
				}
			case 3:
				if t[i] == '1' {
					p4Count++
				}
			case 4:
				if t[i] == '1' {
					p5Count++
				}
			case 5:
				if t[i] == '1' {
					p6Count++
				}
			case 6:
				if t[i] == '1' {
					p7Count++
				}
			case 7:
				if t[i] == '1' {
					p8Count++
				}
			case 8:
				if t[i] == '1' {
					p9Count++
				}
			case 9:
				if t[i] == '1' {
					p10Count++
				}
			case 10:
				if t[i] == '1' {
					p11Count++
				}
			case 11:
				if t[i] == '1' {
					p12Count++
				}
			}
		}
	}

	fmt.Println(p1Count, p2Count, p3Count, p4Count, p5Count)
	gamma := ""
	epsilon := ""

	if p1Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p2Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p3Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p4Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p5Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p6Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p7Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p8Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p9Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}
	if p10Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}

	if p11Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}

	if p12Count > 500 {
		gamma += "1"
		epsilon += "0"
	} else {
		gamma += "0"
		epsilon += "1"
	}

	fmt.Println(gamma, epsilon)

	ei, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	gi, err := strconv.ParseInt(epsilon, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Power Consumption:", ei*gi)
}
