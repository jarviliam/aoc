package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() []string {
	f, err := os.Open("./t.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	out := make([]string, 0)
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func sp(in string) ([]string, []string) {
	c := strings.Split(in, "|")
	return strings.Fields(c[0]), strings.Fields(c[1])
}

func pt1() {
	in := read()
	ans := 0
	for _, cse := range in {
		_, out := sp(cse)
		for _, v := range out {
			if len(v) == 3 || len(v) == 4 || len(v) == 2 || len(v) == 7 {
				ans++
			}
		}
	}
	fmt.Println(ans)

}

/* # 0: abcefg (6)
# 6: abdefg (6)
# 9: abcdfg (6)

# 2: acdeg (5)
# 3: acdfg (5)
# 5: abdfg (5)

# 1: cf (2)
# 4: bcdf (4)
# 7: acf (3)
# 8: abcdefg (7)
*/

func permutations(signals []string) {
	charmap := make(map[string]string)
	cf := ""
	for _, v := range signals {
		if len(v) == 2 {
			//charmap
			cf = v
			break
		}
	}
	for _, v := range signals {
		if len(v) == 6 && (strings.Contains(v, string(cf[0])) != (strings.Contains(v, string(cf[1])))) {
			if strings.Contains(v, string(cf[0])) {
				charmap[string(cf[0])] = "f"
				charmap[string(cf[1])] = "c"
			} else {
				charmap[string(cf[1])] = "f"
				charmap[string(cf[0])] = "c"
			}
		}
	}
	for _, v := range signals {
		if len(v) == 3 { // 7
			for _, c := range v {
				if !strings.Contains(cf, string(c)) {
					charmap[string(c)] = "a"
				}
			}
		}
	}

}

func main() {
	in := read()
	ans := 0
	for _, cse := range in {
		signals, out := sp(cse)
		charmap := make(map[int]string)
		for _, v := range signals {
			if len(signals) == 2 {
				charmap[1] = v
				continue
			}
			if len(signals) == 3 {
				charmap[7] = v
				continue
			}
			if len(signals) == 4 {
				charmap[4] = v
				continue
			}
			if len(signals) == 7 {
				charmap[8] = v
			}
		}
		for i := 1; i <= 9; i++ {

		}
	}
	fmt.Println(ans)

}
