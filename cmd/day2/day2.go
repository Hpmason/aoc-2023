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
	dat, err := os.ReadFile("data/day2.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	lines := strings.Split(text, "\n")

	jobs := make(chan string, len(lines))
	results := make(chan int, len(lines))

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(jobs, results, maxPull)
	}

	skipCount := 0
	// Queue up jobs
	for _, line := range lines {
		// Deal with blank lines
		if len(line) == 0 {
			skipCount += 1
			continue
		}
		jobs <- line
	}
	close(jobs)

	// Reduce results
	result := 0
	for r := 1; r <= len(lines) - skipCount; r++ {
		result += <-results
	}
	fmt.Println("Result: ", result)
}

func calc(lines []string) {
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue;
		}
		game, err := parseGame(line)
		if err != nil {
			panic(err)
		}
		pull := game.getMinPull()
		result += pull.product()
	}
	fmt.Println("Results: ", result)
}

func worker(jobs <-chan string, results chan<- int, maxPull pull) {
	for line := range jobs {
		game, err := parseGame(line)
		if err != nil {
			panic(err)
		}
		pull := game.getMinPull()
		results <- pull.product()
	}
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
				pull.red += int(count)
			case "green":
				pull.green += int(count)
			case "blue":
				pull.blue += int(count)
			default:
				return game, errors.New("Unexpected color: \"" + color + "\"")
			}
		}
		pulls[i] = pull
	}
	game.pulls = pulls
	return game, nil
}

func (g *game) getMinPull() pull {
	minPull := pull {}
	for _, pull := range g.pulls {
		minPull.red = max(minPull.red, pull.red)
		minPull.green = max(minPull.green, pull.green)
		minPull.blue = max(minPull.blue, pull.blue)
	}
	return minPull
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (g *game) isValid(maxPull pull) bool {
	minPull := g.getMinPull()
	if minPull.red > maxPull.red {
		return false
	}
	if minPull.green > maxPull.green {
		return false
	}
	if minPull.blue > maxPull.blue {
		return false
	}
	return true
}

func (p *pull) product() int {
	return p.red * p.green * p.blue
}


