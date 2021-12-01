package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func in(i string) ([]string, error) {
	file, err := os.Open(i)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in, scanner.Err()
}

func pt1(lines []string) int {
	ans := 0
	acc := 0
	seen := make(map[int]struct{})
	pointer := 0
	for {
		if _, ok := seen[pointer]; ok {
			ans = acc
			break
		}
		seen[pointer] = struct{}{}
		words := strings.Fields(lines[pointer])
		if words[0] == "acc" {
			v, _ := strconv.Atoi(words[1])
			acc += v
			pointer += 1
		}
		if words[0] == "nop" {
			pointer += 1
		}
		if words[0] == "jmp" {
			v, _ := strconv.Atoi(words[1])
			pointer += v
		}

	}
	return ans
}

/**
I spend an hour trying to debug,
why the hell it was looping the same values despite modifying list from the lines.
Looking up that golang := copies by value
I copied the array into list. However, for some reason
it referenced the original array.
I even tried the copy function and it wasn't working
until I had to make an interface.
*/
func pt2(lines []string) int {
	ans := 0
	for i := range lines {
		//fmt.Println(lines[2])
		list := make([]string, len(lines))
		copy(list, lines)
		//list := lines
		//copy(list,lines)
		spl := strings.Fields(list[i])
		if spl[0] == "nop" {
			list[i] = "jmp " + spl[1]
		}
		if spl[0] == "jmp" {
			list[i] = "nop " + spl[1]
		}
		if spl[0] == "acc" {
			continue
		}
		itr := 0
		acc := 0
		pointer := 0
		/*if i == 7 {
			fmt.Println(list, i)
			fmt.Println(lines, i)
		}*/
		for (0 <= pointer) && (pointer < len(list)) && itr < 1000 {
			itr += 1
			words := strings.Fields(list[pointer])
			//fmt.Println(list, i)
			if words[0] == "acc" {
				v, _ := strconv.Atoi(words[1])
				acc += v
				pointer += 1
			}
			if words[0] == "nop" {
				pointer += 1
				// fmt.Println(words)
			}
			if words[0] == "jmp" {
				v, _ := strconv.Atoi(words[1])
				pointer += v
			}

		}
		if pointer == len(list) {
			fmt.Println("here")
			ans = acc
		}
	}
	return ans
}

func main() {
	lines, _ := in("input.txt")
	fmt.Println(pt1(lines))
	fmt.Println(pt2(lines))

}
