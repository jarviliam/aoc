package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func countRune(str string, ch rune) int {
	cnt := 0
	for _, i := range str {
		if i == ch {
			cnt++
		}
	}
	return cnt
}

func part1(inputs []string) int {
	//God Bless Regexp
	// Format: 4-5 m : mmmmm
	reg, _ := regexp.Compile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")

	valid := 0
	for _, passport := range inputs {
		match := reg.FindStringSubmatch(passport)

		minVal, _ := strconv.Atoi(match[1])
		maxVal, _ := strconv.Atoi(match[2])
		target := rune(match[3][0])
		cnt := countRune(match[4], target)
		if cnt >= minVal && cnt <= maxVal {
			valid++
		}
	}
	return valid
}

func part2(inputs []string) int {
	reg, _ := regexp.Compile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")
	valid := 0
	for _, passport := range inputs {
		match := reg.FindStringSubmatch(passport)

		minVal, _ := strconv.Atoi(match[1])
		maxVal, _ := strconv.Atoi(match[2])
		target := match[3][0]
		pass := match[4]

		minVal -= 1
		maxVal -= 1
		if maxVal >= len(pass) {
			continue
		}
        //Damn you go for not having XOR
		if ((pass[minVal] == target) || (pass[maxVal] == target)) && !((pass[maxVal] == target) && (pass[minVal] == target)) {
			valid++
		}
	}
	return valid
}

func loadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs, scanner.Err()
}

func main() {
	inputs, _ := loadLines("./input.txt")
	out1 := part1(inputs)
	out2 := part2(inputs)
	log.Printf("Answer 1 : %d\n", out1)
	log.Printf("Answer 2 : %d\n", out2)
}
