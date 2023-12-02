package main

import (
	"fmt"
	"strconv"
	// "io"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("data/day1-ex.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	fmt.Print(text)
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	res := 0
	for _, line := range lines {
		res += parseLine(line)
	}
	fmt.Println("Result: ", res)
}

func parseLine(line string) int {
	d1 := firstDigit(line)
	d2 := lastDigit(line)
	return d1*10 + d2
}

func firstDigit(line string) int {
	strconv.ParseInt(line, 10, 64)
	for _, c := range line {
		if isDigit(c) {
			i, _ := strconv.ParseInt(string(c), 10, 64)
			return int(i)
		}
	}
	return 0
}

func lastDigit(line string) int {
	for i := range line {
		c := line[len(line)-1-i]
		if isDigit(c) {
			i, _ := strconv.ParseInt(string(c), 10, 64)
			return int(i)
		}
	}
	return 0
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
