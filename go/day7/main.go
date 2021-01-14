package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reader(path string) ([]string, error) {
	file, err := os.Open(path)
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

type node struct {
    id string
    children map[string]*node
}

/*
* vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
* faded blue bags contain no other bags.
 */
func pt1(in []string) int {
	ans := 0
	hold := make(map[string]struct{})
	for _, v := range in {
		if v != "" {
			words := strings.Fields(v)
			//holder := words[0] + words[1] + words[2]
			//If it has no bags
			if words[len(words)-3] == "no" {
				continue
			} else {
				idx := 4
				for {
					if idx == len(words) {
						break
					}
					bag := words[idx] + words[idx+1] + words[idx+2] + words[idx+3]
					if bag[len(bag)-1] == '.' {
						bag = bag[:len(bag)-1]
						fmt.Println(bag)
					}
					if bag[len(bag)-1] == ',' {
						bag = bag[:len(bag)-1]
						fmt.Println(bag)
					}
					if bag[len(bag)-1] == 's' {
						bag = bag[:len(bag)-1]
						fmt.Println(bag)
					}
					if hold[bag] != struct{}{} {
                        hold[bag] =
						//bag = bag[:len(bag)-1]
						//fmt.Println(bag)
					}
					hold[bag] = struct{}{}
				}
			}
		}
	}
	return ans
}
func pt2(in []string) int {
	ans := 0

	return ans
}

func main() {

	in, _ := reader("test.txt")
	ans1 := pt1(in)
	ans2 := pt2(in)
	fmt.Println(ans1)
	fmt.Println(ans2)
}
