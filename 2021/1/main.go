package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read() []int {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	o := make([]int, 0)
	for s.Scan() {
		v, _ := strconv.Atoi(s.Text())
		o = append(o, v)
	}
	return o
}
func solve1() {
	input := read()
	prev := input[0]
	inc := 0
	for i, v := range input {
		if i == 0 {
			continue
		}
		if v > prev {
			inc++
		}
		prev = v
	}
	fmt.Printf("Sol 1 Out: %d\n", inc)
}
func solve2() {
	in := read()
	cmp := func(a []int, b []int) bool {
		sumA := 0
		for _, v := range a {
			sumA += v
		}
		sumB := 0
		for _, v := range b {
			sumB += v
		}
		return sumB > sumA
	}
	idx := 0
	inc := 0
	for {
		if idx+3 >= len(in) {
			break
		}
		a := in[idx : idx+3]
		b := in[idx+1 : idx+4]
		out := cmp(a, b)
		if out {
			inc++
		}
		//fmt.Printf("Case : %d ; %v\n",idx,out)
		idx++
	}
	fmt.Printf("Sol 2 Out: %d\n", inc)
}

func main() {
	solve1()
	solve2()
}
