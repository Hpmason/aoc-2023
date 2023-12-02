package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const numWorkers = 10
	dat, err := os.ReadFile("data/day1.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	lines := strings.Split(text, "\n")
	jobs := make(chan string, len(lines))
	results := make(chan int, len(lines))


	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(jobs, results)
	}
	// Queue up jobs
	for _, line := range lines {
		jobs <- line
	}
	close(jobs)

	// Reduce results
	result := 0
	for r := 1; r <= len(lines); r++ {
		result += <-results
	}
	fmt.Println("Result: ", result)
}

func worker(jobs <-chan string, results chan<- int) {
	for j := range jobs {
		results <- parseLine(j)
	}
}

func parseLine(line string) int {
	d1 := firstDigit(line)
	d2 := lastDigit(line)
	return d1*10 + d2
}

func firstDigit(line string) int {
	strconv.ParseInt(line, 10, 64)
	for i, c := range line {
		if isDigit(c) {
			i, _ := strconv.ParseInt(string(c), 10, 64)
			return int(i)
		}
		i, found := hasDigitWord(line, i)
		if found {
			return i
		}
	}
	return 0
}

func lastDigit(line string) int {
	for i := range line {
		i = len(line)-1-i
		c := line[i]
		if isDigit(c) {
			num, _ := strconv.ParseInt(string(c), 10, 64)
			return int(num)
		}
		num, found := hasDigitWord(line, i)
		if found {
			return num
		}
	}
	return 0
}


var digits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func hasDigitWord(line string, start int) (int, bool) {
	slice := line[start:]
	for i, word := range digits {
		if len(slice) < len(word) {
			continue
		}
		if strings.EqualFold(slice[:len(word)], word) {
			return i+1, true
		}
	}
	return 0, false
}


func isDigit(c interface {}) bool {
	switch c.(type) {
	case byte:
		c := c.(byte)
		return '0' <= c && c <= '9'
	case rune:
		c := c.(rune)
		return '0' <= c && c <= '9'
	default:
		return false
	}
}
