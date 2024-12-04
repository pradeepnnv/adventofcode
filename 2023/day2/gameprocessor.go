package day2

import (
	"log"
	"strconv"
	"strings"
)

func GameProcessorDay1(input []string) int {
	sum := 0
	for _, str := range input { // game
		// gameId := 0
		// games := []string{}
		split1 := strings.Split(str, ":")
		// fmt.Println(split1[0])
		// fmt.Println(strings.Split(split1[0], " ")[1])
		gameId, err := strconv.Atoi(strings.Split(split1[0], " ")[1])
		if err != nil {
			log.Fatalf("%s", err)
		}
		games := strings.Split(split1[1], ";")
		// fmt.Println(games)
		possibleGame := true
		for _, game := range games { //draw
			// fmt.Println("processing ", game)
			split2 := strings.Split(game, ", ")
			// fmt.Println(split2)

			for _, c := range split2 { //color
				split3 := strings.Split(strings.TrimSpace(c), " ")
				// fmt.Println("Split3 is :", split3)
				n, err := strconv.Atoi(split3[0])
				if err != nil {
					log.Fatalf("%s", err)
				}
				red := 0
				blue := 0
				green := 0
				switch split3[1] {
				case "red":
					red += n
				case "blue":
					blue += n
				case "green":
					green += n
				}
				// fmt.Printf("Final tally for game %d is red:%d, blue:%d, green:%d\n", gameId, red, blue, green)

				if red > 12 || green > 13 || blue > 14 {
					// fmt.Println("This game is not possible")
					possibleGame = false
					break
				}
			}

		}
		if possibleGame {
			sum += gameId
		}

	}
	return sum
}

func GameProcessorPart2(games []string) int {
	sum := 0
	for _, game := range games {
		split1 := strings.Split(game, ":")
		red := 0
		green := 0
		blue := 0
		draws := strings.Split(split1[1], ";")
		fDraw := true // indicates first draw of a game
		for _, draw := range draws {
			colors := strings.Split(strings.TrimSpace(draw), ", ")
			// fmt.Println("Colors is :", colors)
			for _, c := range colors {
				split3 := strings.Split(strings.TrimSpace(c), " ")
				n, err := strconv.Atoi(split3[0])
				// fmt.Println("Value of n is", n)
				if err != nil {
					log.Fatalf("%s", err)
				}

				switch split3[1] {
				case "red":
					if fDraw || red == 0 || n > red {
						red = n
					}
				case "blue":
					if fDraw || blue == 0 || n > blue {
						blue = n
					}
				case "green":
					if fDraw || green == 0 || n > green {
						green = n
					}
				}
			}
			fDraw = false
			// fmt.Printf("red balls : %d, green balls : %d, blue balls; %d\n", red, green, blue)
		}
		// fmt.Printf("Minimum red balls : %d, green balls : %d, blue balls; %d for the game ==>%s\n", red, green, blue, game)
		sum += red * green * blue
	}
	return sum
}
