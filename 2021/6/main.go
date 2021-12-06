package main

import (
	"bufio"
	"fmt"
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

func main() {
	in := read()
	days := 256
	counter := make([]int, 10)
	for _, v := range in {
		counter[v] += 1
	}
	for i := 0; i < days; i++ {
		nv := make([]int, 10)
		for j := range counter {
			if j == 0 {
				continue
			}
			nv[j-1] += counter[j]
		}
		nv[6] += counter[0]
		nv[8] += counter[0]
		counter = nv
	}
	sum := 0
	for _, v := range counter {
		sum += v
	}
	fmt.Println(sum)
	fmt.Printf("%v\n", in)
}
