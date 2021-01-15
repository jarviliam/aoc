package main

import (
	"bufio"
	"fmt"
	"github.com/tebeka/deque"
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
	id       string
	children map[string]*node
}

/*
* vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
* faded blue bags contain no other bags.
 */
func pt1(in []string) int {
	ans := 0
	hold := make(map[string][]string)
	for _, v := range in {
		if v != "" {
			words := strings.Fields(v)
			holder := words[0] + words[1] + words[2]
			holder = holder[:len(holder)-1]
			fmt.Println(holder)
			//If it has no bags
			if words[len(words)-3] == "no" {
				continue
			} else {
				idx := 4
				for {
					if idx >= len(words) {
						break
					}
					bag := words[idx] + words[idx+1] + words[idx+2] + words[idx+3]
					if bag[len(bag)-1] == '.' {
						bag = bag[:len(bag)-1]
					}
					if bag[len(bag)-1] == ',' {
						bag = bag[:len(bag)-1]
					}
					if bag[len(bag)-1] == 's' {
						bag = bag[:len(bag)-1]
					}
					if bag[0] >= '0' && bag[0] <= '9' {
						fmt.Println(bag, "to ", bag[1:])
						bag = bag[1:]
					}
					if _, ok := hold[bag]; !ok {
						hold[bag] = []string{}
					}
					hold[bag] = append(hold[bag], holder)
					//fmt.Println(bag)
					idx += 4
				}
			}
		}
	}
	for k, v := range hold {
		fmt.Println(k, "value is ", v)
	}
	var exists = struct{}{}

    /*
    * BFS
    */
	seen := make(map[string]struct{})
	target := "shinygoldbag"
	q := deque.New()
    //Append the Target
	q.Append(target)
	for {
        //Break when Empty
		if q.Len() == 0 {
			break
		}
        //Get Left Value
		x, _ := q.PopLeft()
        //Convert Interface -> String
		z := fmt.Sprintf("%v", x)

        //If it Exists within our Seen Set. Skip itr
		if _, ok := seen[z]; ok {
			continue
		}
        //Append
		seen[z] = exists
        //For
		for _, y := range hold[z] {
			q.Append(y)
		}
	}
	ans = len(seen) - 1
	return ans
}
func pt2(in []string) int {
	ans := 0

	return ans
}

func main() {

	in, _ := reader("input.txt")
	ans1 := pt1(in)
	ans2 := pt2(in)
	fmt.Println(ans1)
	fmt.Println(ans2)
}
