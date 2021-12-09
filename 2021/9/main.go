package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func read() [][]int {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	out := make([][]int, 0)
	for s.Scan() {
		line := s.Text()
		l := make([]int, 0)
		for _, v := range line {
			i, _ := strconv.Atoi(string(v))
			l = append(l, i)
		}
		out = append(out, l)
	}
	return out
}

func conv(in byte) int {
	c, _ := strconv.Atoi(string(in))
	return c
}

func pt1() {
	in := read()
	sum := 0
	rC := len(in)
	cC := len(in[0])
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	for i, v := range in {
		for j, w := range v {
			lowest := true
			for k := 0; k < 4; k++ {
				rr := i + dr[k]
				cc := j + dc[k]
				if 0 <= rr && rr < rC && 0 <= cc && cc < cC && in[rr][cc] <= w {
					lowest = false
				}

			}
			if lowest {
				sum += w + 1
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	//This was stupid
	pt1()

	in := read()
	seen := make(map[string]interface{})
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	s := make([]int, 0)
	for i, v := range in {
		for j, w := range v {
			see := fmt.Sprintf("%d-%d", i, j)
			_, o := seen[see]
			if !o && w != 9 {
				size := 0
				queue := make([][]int, 0)
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					p := queue[0]
					queue = queue[1:]
                    i = p[0]
                    j = p[1]
					_, ok := seen[fmt.Sprintf("%d-%d", p[0], p[1])]
					fmt.Println(p[0],p[1],ok)
					if ok {
						continue
					}
					fmt.Println(i,j)
					seen[fmt.Sprintf("%d-%d", i, j)] = -1
					fmt.Println(seen)
					size += 1
					for k := 0; k < 4; k++ {
						rr := i + dr[k]
						cc := j + dc[k]
						if 0 <= rr && rr < len(in) && 0 <= cc && cc < len(v) && in[rr][cc] != 9 {
							queue = append(queue, []int{rr, cc})
						}
					}
				}
				s = append(s, size)
			}
		}
	}
	sort.Ints(s[:])
	fmt.Println(s[len(s)-3] * s[len(s)-2] * s[len(s)-1])
}
