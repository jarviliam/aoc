package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read() {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("File ")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	x, y := 0, 0
	for s.Scan() {
		i := strings.Split(s.Text(), " ")
		if len(i) != 2 {
			panic("Input Miss")
		}
		v, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "up":
			y += v * -1
		case "down":
			y += v
		case "forward":
			x += v
		}
	}
	fmt.Println(x * y)
}
func read2() {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("File")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	x, y, aim := 0, 0, 0
	for s.Scan() {
		i := strings.Split(s.Text(), " ")
		if len(i) != 2 {
			panic("Input Miss")
		}
		v, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "up":
			aim += v * -1
		case "down":
			aim += v
		case "forward":
			x += v
			y += v * aim
		}
	}
	fmt.Println(x * y)
}

func main() {
	read()
	read2()
}
