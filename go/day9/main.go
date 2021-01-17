package main

import (
	"bufio"
	"fmt"
	"github.com/mxschmitt/golang-combinations"
	"os"
	"strconv"
	//"strings"
)

func read(i string) ([]string, error) {
	file, err := os.Open(i)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//v, _ := strconv.Atoi(scanner.Text())
		in = append(in, scanner.Text())
	}
	return in, scanner.Err()
}

func pt1(x []string) uint64 {
	thresh := 5
	ans := uint64(0)

	for i := range x[thresh:] {
		i += thresh
		ok := true
		prev := x[i-thresh : i]
		test := combinations.Combinations(prev, 2)
		for _, v := range test {
			a, _ := strconv.Atoi(v[0])
			b, _ := strconv.Atoi(v[1])
			target, _ := strconv.Atoi(x[i])
			if (a + b) == target {
				ok = false
			}
		}
		if ok {
			out, _ := strconv.Atoi(x[i])
			ans = out
			break
		}
	}

	return ans
}
func min(y []uint64) uint64 {
	min := y[0]
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}
func max(y []uint64) uint64 {
	max := y[0]
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
func sum(arr []uint64) uint64 {
	res := uint64(0)
	for _, v := range arr {
		res += v
	}
	return res
}
func conv(arr []string) []uint64 {
	x := make([]uint64, len(arr))
	for _, v := range arr {
		y, _ := strconv.Atoi(v)
		x = append(x, uint64(y))
	}
	return x
}
func pt2(x []string, target uint64) uint64 {
	ans := uint64(0)
	y := conv(x)
	for i := range y {
		for k, _ := range y[i+1:] {
			k += i
			if sum(y[i:k]) == target && ans == 0 {
				ans = min(y[i:k]) + max(y[i:k])
				fmt.Println(min(y[i:k]))
				fmt.Println(max(y[i:k]))
			}
		}
	}

	return ans
}

func main() {

	in, _ := read("input.txt")

	p1 := pt1(in)
	fmt.Println(p1)
	fmt.Println(pt2(in, p1))
}
