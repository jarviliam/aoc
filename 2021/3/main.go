package main

import (
	"bufio"
	"fmt"
	"os"
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
func p1() {
	in := read()
	maxX := len(in[0]) - 1
	maxY := len(in)
	gamma := 0
	epilson := 0
	for x := 0; x <= maxX; x++ {
		count := 0
		for _, v := range in {
			if v[x] == '1' {
				count += 1
			}
		}
		gamma = gamma << 1
		epilson = epilson << 1
		if count >= maxY/2 {
			gamma = gamma | (gamma + 1)
		} else {
			epilson = epilson | (epilson + 1)
		}
	}
	fmt.Println(epilson * gamma)
}

func oxygen(list []string, bit int) string {
	if len(list) == 1 {
		return list[0]
	}
	ones := make([]string, 0)
	zeroes := make([]string, 0)

	for _, v := range list {
		if v[bit] == '1' {
			ones = append(ones, v)
		} else {
			zeroes = append(zeroes, v)
		}
	}
	if len(ones) >= len(zeroes) {
		return oxygen(ones, bit+1)
	} else {
		return oxygen(zeroes, bit+1)
	}
}
func co2(list []string, bit int) string {
	if len(list) == 1 {
		return list[0]
	}
	ones := make([]string, 0)
	zeroes := make([]string, 0)

	for _, v := range list {
		if v[bit] == '1' {
			ones = append(ones, v)
		} else {
			zeroes = append(zeroes, v)
		}
	}
	if len(zeroes) <= len(ones) {
		return co2(zeroes, bit+1)
	} else {
		return co2(ones, bit+1)
	}
}
func formt(in string) int {
	res := 0
	for _, v := range in {
		if v == '1' {
			res += 1
		}
		res *= 2
	}
	return res / 2
}

func main() {

	in := read()

	fmt.Println(formt(oxygen(in, 0)) * formt(co2(in, 0)))
}
