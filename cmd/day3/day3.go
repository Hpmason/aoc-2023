package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	dat, err := os.ReadFile("data/day3.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	fmt.Println(text)
	lines := strings.Split(text, "\n")

	// Collection of starting positions of part numbers mapped to its length
	partNumbers := make(map[point]int)
	gears := make([]point, 0)
	for j, line := range lines {
		for i := 0; i < len(line); i++ {
			if isNum(line[i]) {
				start := point{i, j}
				for {
					if i+1 < len(line) && isNum(line[i+1]) {
						i++
					} else {
						break
					}
				}
				partNumbers[start] = i - start.x + 1
				if partNumbers[start] == 0 {
					fmt.Println("Pos ", start.x, ", ", start.y)
				}
			} 
			if isGear(byte(line[i])) {
				gears = append(gears, point{i, j})
			}
		}
	}
	result := 0
	foundNumbers := make([]point, 0)
	for _, p := range gears {
		for j := p.y - 1; j <= p.y + 1; j++ {
			for i := p.x - 1; i <= p.x + 1; i++ {
				if isNum(lines[j][i]) {
					k := i
					for {
						if k == 0 || !isNum(lines[j][k-1]) {
							break
						} else {
							k--
						} 
					}
					key := point{k, j}
					if _, prs := partNumbers[key]; prs && !slices.Contains(foundNumbers, key) {
						foundNumbers = append(foundNumbers, key)
					}
				}
			}
		}
		if len(foundNumbers) == 2 {
			product := -1
			for _, n := range foundNumbers {
				length := partNumbers[n]
				num, err := strconv.ParseInt(lines[n.y][n.x:n.x+length], 10, 64)
				if err != nil {
					panic(fmt.Errorf("\"%s\": x = %d, len = %d", lines[n.y][n.x:], n.x, length))
				}
				
				if product == -1 {
					product = int(num)
				} else {
					product *= int(num)
				}
			}
			result += product
		}
		foundNumbers = foundNumbers[:0]
	}
	fmt.Println(partNumbers)
	fmt.Println(gears)
	fmt.Println("Result: ", result)
}

func isGear(c byte) bool {
	return c == '*'
}

func isSymbol(c byte) bool {
	return c != '.' && !isNum(byte(c))
}

func isNum(c byte) bool {
	return '0' <= c && c <= '9'
}
