package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			inputs = append(inputs, scanner.Text())
		}
	}
	return inputs, scanner.Err()
}

func treeSmashes(list []string, r int, d int) int {
	answer := 0
	y := 0
	x := 0
	height := len(list)
	length := len(list[0])

	for y < height-d {
		x = x + r
		if x >= length {
			x = x % length
		}
		y = y + d
		if list[y][x] == '#' {
			answer = answer + 1
		}
	}
	return answer
}

func main() {
	inputs, _ := inputLines("./input.txt")
	ans1 := treeSmashes(inputs, 3, 1)
	fmt.Printf("Answer : %d\n", ans1)
}
