package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func load(in string) ([]string, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, scanner.Err()
}

func convert(in []string) []int {
	out := []int{0}
	for _, v := range in {
		y, _ := strconv.Atoi(v)
		out = append(out, y)
	}
	return out
}

func pt1(in []int) int {
	ans := 0
	sort.Ints(in)
	in = append(in, in[len(in)-1]+3)
	diff1, diff3 := 0, 0

	for i, _ := range in {
		if i == len(in)-1 {
			continue
		}
		diff := in[i+1] - in[i]
		if diff == 1 {
			diff1 += 1
		}
		if diff == 3 {
			diff3 += 1
		}
	}
	ans = diff1 * diff3
	return ans
}

func makeRange(x int, y int) []int {
	var out []int
	for x < y {
		out = append(out, x)
		x++
	}
	return out
}

func dyanimcProgramming(sorted []int, i int, seen map[int]int) int {
	if i == len(sorted)-1 {
		return 1
	}
	if _, ok := seen[i]; ok {
		return seen[i]
	}
	answer := 0
	for _, k := range makeRange(i+1, len(sorted)) {
		if sorted[k]-sorted[i] <= 3 {
			answer += dyanimcProgramming(sorted, k, seen)
		}
	}
	seen[i] = answer
	return answer
}

func pt2(in []int) int {
	seen := make(map[int]int)
	return dyanimcProgramming(in, 0, seen)
}

func main() {

	input, _ := load("test.txt")

	fmt.Println(pt1(convert(input)))
	fmt.Println(pt2(convert(input)))

}
