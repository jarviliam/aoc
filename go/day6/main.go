package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reader(path string) ([]string, error) {
	file, er := os.Open(path)

	if er != nil {
		return nil, er
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs, scanner.Err()
}

/*
* Arrange Set, On New Line increment set length
 */
func pt1(list []string) int {
	ans := 0
	set := make(map[rune]struct{})
	var exist = struct{}{}

	for _, l := range list {
		if l == "" {
			ans += len(set)
			set = make(map[rune]struct{})
		} else {
			for _, v := range l {
				set[v] = exist
			}
		}
	}
	return ans
}

func makeSet() map[rune]struct{} {
	ascii := "abcdefghijklmnopqrstuvwxyz"
	set := make(map[rune]struct{})
	var exist = struct{}{}
	for _, c := range ascii {
		set[c] = exist
	}
	return set
}

/*
 * Make Set of A - Z Runes
 * If the Line does not contain the char and it exists in set
 * Delete it from the entry
 *
 */
func pt2(list []string) int {
	ans := 0
	ascii := "abcdefghijklmnopqrstuvwxyz"
	set := makeSet()
	for _, l := range list {
		if l == "" {
			ans += len(set)
			set = makeSet()
		} else {
			//Loop ascii
			for _, v := range ascii {
				if !strings.ContainsRune(l, v) && set[v] == struct{}{} {
					delete(set, v)
				}
			}
		}
	}
	return ans
}

func main() {
	input, _ := reader("input.txt")
	part1 := pt2(input)

	fmt.Println(part1)
}
