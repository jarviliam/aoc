package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func read() []int {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	out := make([]int, 0)
	for _, v := range strings.Split(s.Text(), ",") {
		conv, _ := strconv.Atoi(v)
		out = append(out, conv)
	}
	return out
}

func high(in []int) int {
	high := 0
	for _, v := range in {
		if v > high {
			high = v
		}
	}
	return high
}
func pt1() {
	in := read()
	y := high(in)
	cheapest := 100000000
	for i := 0; i <= y; i++ {
		spot := 0
		for _, v := range in {
			spot += int(math.Abs(float64(v - i)))
		}
		if spot < cheapest {
			cheapest = spot
		}
	}
	fmt.Printf("Part 1 : %d\n", cheapest)
}
func pt2() {
	in := read()
	y := high(in)
	cheapest := 100000000
	for i := 0; i <= y; i++ {
		spot := 0
		for _, v := range in {
			diff := math.Abs(float64(v - i))
			spot += int((diff * (diff + 1)) / 2)
		}
		if spot < cheapest {
			cheapest = spot
		}
	}
	fmt.Printf("Part 2 : %d\n", cheapest)
}

func main() {
	pt1()
	pt2()
}
