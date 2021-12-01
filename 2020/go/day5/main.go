package main

import (
	"bufio"
	"fmt"
	"os"
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
		if scanner.Text() != "" {
			inputs = append(inputs, scanner.Text())
		}
	}
	return inputs, scanner.Err()
}

//Binary Implementation
//Each char represents a bit position
//So Division by 2 as you go right will get you the bit val
// B = 1 F = 0
func part1(inputs []string) int {
	ans := 0
	ids := make(map[int]struct{})
	var exist = struct{}{}
	for _, l := range inputs {
		row, col := 0, 0
		rp := 64
		cp := 4

		for _, c := range l {
			if c == 'B' {
				row += rp
				rp /= 2
			} else if c == 'F' {
				rp /= 2
			}

			if c == 'R' {
				col += cp
				cp /= 2
			} else if c == 'L' {
				cp /= 2
			}
		}
		seat_id := row*8 + col
        //fmt.Printf("Answer :%d",seat_id)
		ids[seat_id] = exist
       if(ans < seat_id){
           ans = seat_id
       }else{

       }
	}
	return ans
}

func main() {
	in, _ := reader("./input.txt")
    ans := part1(in)
    fmt.Printf("Answer :%d",ans)
}
