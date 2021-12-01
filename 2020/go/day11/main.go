package main

import (
	"bufio"
	"fmt"
	"os"
)

func load(name string) ([]string, error) {
	file, err := os.File(name)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	var input []string
	Scanner := bufio.NewScanner(file)

	for Scanner.Scan() {
		input = append(input, Scanner.Text())
	}

	return input, nil

}

func pt1(input []string) int {

    return 0
}

func main() {

    inputs, _:= load("test.txt")
    fmt.Println(pt1(inputs))
}
