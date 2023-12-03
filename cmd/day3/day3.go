package main

import (
	"fmt"
	"os"
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
	symbols := make([]point, 0)
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
			} 
			if isSymbol(byte(line[i])) {
				symbols = append(symbols, point{i, j})
			}
		}
	}
	result := 0
	for _, p := range symbols {
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
					length, prs := partNumbers[key]
					if !prs {
						panic(fmt.Errorf("Could not get number at %d, %d", k, j))
					}
					if length > 0 {
						num, err := strconv.ParseInt(lines[j][k:k+length], 10, 64)
						if err != nil {
							panic(err)
						}
						result += int(num)
						partNumbers[key] = 0
					}
				}
			}
		} 
	}
	fmt.Println(partNumbers)
	fmt.Println(symbols)
	fmt.Println("Result: ", result)
}

func isSymbol(c byte) bool {
	return c != '.' && !isNum(byte(c))
}

func isNum(c byte) bool {
	return '0' <= c && c <= '9'
}
