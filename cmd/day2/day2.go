package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const numWorkers = 10
	var maxPull = pull{12, 13, 14}
	dat, err := os.ReadFile("data/day2-ex1.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	lines := strings.Split(text, "\n")
	// for _, line := range lines {
	// 	if len(line) == 0 {
	// 		continue
	// 	}
	// 	game, err := parseGame(line)
	// 	if err != nil {
	// 		for i := range line {
	// 			fmt.Println(line[i])
	// 		}
	// 		fmt.Println("Error parsing: ", err)
	// 	} else {
	// 		fmt.Println(game)
	// 	}
	// }
	
	
	jobs := make(chan string, len(lines))
	results := make(chan int, len(lines))

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(jobs, results, maxPull)
	}
	// Queue up jobs
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		jobs <- line
	}
	close(jobs)

	// Reduce results
	result := 0
	for r := 1; r <= len(lines) - 1; r++ {
		result += <-results
	}
	fmt.Println("Result: ", result)
}

type game struct {
	num int
	pulls []pull
}

type pull struct {
	red, green, blue int
}


func parseGame(line string) (game, error) {
	game := game{}
	gamePart, pullPart, found := strings.Cut(line, ": ")
	if !found {
		return game, errors.New("\": \" line seperator not found")
	}
	_, numText, found := strings.Cut(gamePart, "Game ")
	if !found {
		return game, errors.New("Could not find \"Game \" prefix")
	}
	num, err := strconv.ParseInt(numText, 10, 64)
	if err != nil {
		return game, errors.New("Could not parse game number")
	}
	game.num = int(num)
	
	pullsText := strings.Split(pullPart, "; ")
	pulls := make([]pull, len(pullsText))
	for i, pullText := range pullsText {
		pullParts := strings.Split(pullText, ", ")
		pull := pull{}
		for _, part := range pullParts {
			countTxt, color, found := strings.Cut(part, " ")
			if !found {
				return game, errors.New("Could find space between count and color")
			}
			count, err := strconv.ParseInt(countTxt, 10, 64)
			if err != nil {
				return game, errors.New("Could not parse count")
			}
			switch color {
			case "red":
				pull.red = int(count)
			case "green":
				pull.green = int(count)
			case "blue":
				pull.blue = int(count)
			}
		}
		pulls[i] = pull
	}
	game.pulls = pulls
	return game, nil
}

func (g *game) isValid(maxPull pull) bool {
	cumulativePull := pull {}
	for _, pull := range g.pulls {
		if pull.red > cumulativePull.red {
			cumulativePull.red = pull.red
		}
		if pull.green > cumulativePull.green {
			cumulativePull.green = pull.green
		}
		if pull.blue > cumulativePull.blue {
			cumulativePull.blue = pull.blue
		}
	}

	if cumulativePull.red > maxPull.red {
		return false
	}
	if cumulativePull.green > maxPull.green {
		return false
	}
	if cumulativePull.blue > maxPull.blue {
		return false
	}
	return true
}

func worker(jobs <-chan string, results chan<- int, maxPull pull) {
	for line := range jobs {
		game, err := parseGame(line)
		if err != nil {
			panic(err)
		}
		fmt.Println(game)
		if game.isValid(maxPull) {
			results <- game.num
		} else {
			results <- 0
		}
	}
}
