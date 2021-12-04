package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Would have been better if I made it a int array but...
func read() ([]string, [][][]string) {
	f, err := os.Open("./i1.txt")
	if err != nil {
		panic("No File")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	o := make([][]string, 0)
	s.Scan()
	called := strings.Split(s.Text(), ",")
	boards := make([][][]string, 0)
	s.Scan()
	for s.Scan() {
		if s.Text() == "" {
			boards = append(boards, o)
			o = make([][]string, 0)
			continue
		}
		o = append(o, strings.Fields(s.Text()))
	}
	boards = append(boards, o)
	fmt.Println(len(boards))
	return called, boards
}

//Verifies Rows and Columns. Marked == ""
func checkBingo(board [][]string, r, c int) bool {
	cBingo := true
	hBingo := true
	//Column Check
	for _, v := range board {
		if v[c] != "" {
			cBingo = false
			break
		}

	}
	//Row Check
	for _, v := range board[r] {
		if v != "" {
			hBingo = false
			break
		}
	}
	return cBingo || hBingo
}

//calculates sum of board
func sum(board [][]string) int {
	sum := 0
	for _, r := range board {
		for _, c := range r {
			if c != "" {
				v, _ := strconv.Atoi(c)
				sum += v
			}
		}
	}
	return sum
}

//Returns the Summation and the remaining amount of marks
func play(board [][]string, mark []string) (int, int) {
	if len(mark) == 0 {
		return -1, -1
	}
	next := mark[0]
	fmt.Printf("Searching for : %s\n", next)
	for ir, r := range board {
		for ic, c := range r {
			//Found Mark
			if c == next {
				board[ir][ic] = ""
				if checkBingo(board, ir, ic) {
					fmt.Println("Bingo'd")
					mul, _ := strconv.Atoi(next)
					return sum(board) * mul, len(mark[1:])
				}
			}
		}
	}
	return play(board, mark[1:])
}

//Returns fastest speed idx
func getFastestTimeIdx(speeds []int) int {
	fastestIdx := -1
	time := -1
	for i, v := range speeds {
		if v > time {
			fastestIdx = i
			time = v
		}
	}
	return fastestIdx
}

// Returns slowest time idx
func getSlowestTimeIdx(speeds []int) int {
	slowestIdx := -1
	slowestTime := 99999
	for i, v := range speeds {
		if v < slowestTime && v > 0 {
			slowestIdx = i
			slowestTime = v
		}
	}
	return slowestIdx
}

func main() {
	marked, boards := read()
	play(boards[0], marked)
	results := make([]int, 0)
	speed := make([]int, 0)

	//Each board calculate Summation and Speed it took. Speed = Len(marks)
	for _, v := range boards {
		out, remain := play(v[:], marked)
		results = append(results, out)
		speed = append(speed, remain)
	}
	fastestIdx := getFastestTimeIdx(speed)
	fmt.Printf("Fastest : %d ; Val : %d", fastestIdx, results[fastestIdx])

	slowestIdx := getSlowestTimeIdx(speed)
	fmt.Printf("Slowest : %d ; Val : %d", slowestIdx, results[slowestIdx])
}
