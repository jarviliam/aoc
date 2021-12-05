package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read() []string {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	o := make([]string, 0)
	for s.Scan() {
		o = append(o, s.Text())
	}
	return o
}

func getCoords(in string) (x1, y1, x2, y2 int) {
	sep := strings.Fields(in)
	orig := strings.Split(sep[0], ",")
	x1, _ = strconv.Atoi(orig[0])
	y1, _ = strconv.Atoi(orig[1])
	after := strings.Split(sep[2], ",")
	x2, _ = strconv.Atoi(after[0])
	y2, _ = strconv.Atoi(after[1])
	return x1, y1, x2, y2
}
//I did this pretty badly
func p1() {
	unique := make(map[string]int, 0)
	in := read()
	for _, v := range in {
		x1, y1, x2, y2 := getCoords(v)
		if x1 != x2 && y1 != y2 {
			continue
		}
		//fmt.Printf("Coords: {%d,%d} : {%d,%d}\n", x1, y1, x2, y2)
		//Go left
		if x2 < x1 {
			for i := x1; i >= x2; i-- {
				//Go up
				if y1 < y2 {
					for j := y1; j <= y2; j++ {
						unique[fmt.Sprintf("%d,%d", i, j)]++
					}
				} else {
					for j := y1; j >= y2; j-- {
						unique[fmt.Sprintf("%d,%d", i, j)]++
					}
				}
			}
		} else {
			for i := x1; i <= x2; i++ {
				//Go up
				if y1 < y2 {
					for j := y1; j <= y2; j++ {
						unique[fmt.Sprintf("%d,%d", i, j)]++
					}
				} else {
					for j := y1; j >= y2; j-- {
						//fmt.Println(j)
						unique[fmt.Sprintf("%d,%d", i, j)]++
					}
				}
			}
		}
	}
	fmt.Printf("Output : %v\n", unique)
	out := 0
	for _, v := range unique {
		if v >= 2 {
			out++
		}
	}
	fmt.Println(out)
}
//Shoulda done it this way
func p2() {
	unique := make(map[string]int, 0)
	in := read()
	for _, v := range in {
		x1, y1, x2, y2 := getCoords(v)
		i, j := x1, y1
		for {
			unique[fmt.Sprintf("%d,%d", i, j)]++
			if i == x2 && j == y2 {
				break
			}
			if i > x2 {
				i--
			} else if i < x2 {
				i++
			}
			if j > y2 {
				j--
			} else if j < y2 {
				j++
			}
		}
	}

	fmt.Printf("Output : %v\n", unique)
	out := 0
	for _, v := range unique {
		if v >= 2 {
			out++
		}
	}
	fmt.Println(out)

}

func main() {
	p1()
    p2()

}
