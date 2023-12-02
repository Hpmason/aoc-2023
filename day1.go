package main

import (
	"fmt"
	"strconv"
	// "io"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("data/day1-ex2.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	res := 0
	for _, line := range lines {
		res += parseLine(line)
	}
	fmt.Println("Result: ", res)
}

func parseLine(line string) int {
	fmt.Println(line)
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
			fmt.Println(i, string(rune(c)))
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
			fmt.Println(i, string(rune(c)))
			return num
		}
	}
	return 0
}


var digits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
const MAX_DIGIT_LEN int = len("three")

func hasDigitWord(line string, start int) (int, bool) {
	slice := line[start:]
	if len(slice) > MAX_DIGIT_LEN {
		slice = slice[:MAX_DIGIT_LEN]
	}
	for i, word := range digits {
		if strings.Contains(slice, word) {
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
