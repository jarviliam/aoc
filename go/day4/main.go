package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		if scanner.Text() != "" {
			inputs = append(inputs, scanner.Text())
		}
	}
	return inputs, scanner.Err()
}

func isInRange(x int, l int, h int) bool { return x >= l && x <= h }
func isFourDigits(x string) bool         { return len(x) == 4 }

var ppscanner = map[string]func(string) bool{
	"byr": func(x string) bool {
		value, err := strconv.Atoi(x)
		if err != nil {
			return false
		}
		return (isFourDigits(x) && isInRange(value, 1920, 2002))
	},
	"iyr": func(x string) bool {
		value, err := strconv.Atoi(x)
		if err != nil {
			return false
		}
		return (isFourDigits(x) && isInRange(value, 2010, 2020))
	},
	"eyt": func(x string) bool {
		value, err := strconv.Atoi(x)
		if err != nil {
			return false
		}
		return (isFourDigits(x) && isInRange(value, 2020, 2030))
	},
	"hgt": func(x string) bool {
		//Get in or CM
		cm := strings.Index(x, "cm")
		in := strings.Index(x, "in")
		if cm == -1 && in == -1 {
			fmt.Println("Not found")
			return false
		}
		if in != -1 {
			inch_value, err := strconv.Atoi(x[:in])
			if err != nil {
				return false
			}
			return inch_value >= 59 && inch_value <= 76
		} else {
			cm_value, err := strconv.Atoi(x[:in])
			if err != nil {
				return false
			}
			return cm_value >= 150 && cm_value <= 76
		}
	},
	"hcl": func(x string) bool {
		if len(x) != 7 {
			return false
		}
		//check hex
		if x[0] != '#' {
			return false
		}
		reg, err := regexp.Compile(`[0-9a-f]*`)
		if err != nil {
			return false
		}
		if !reg.MatchString(x[1:]) {
			return false
		}
		return true
	}, "ecl": func(x string) bool {
		eyemap := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
		if ok := eyemap[x]; ok {
			return true
		}
		return false
	}, "pid": func(x string) bool {
		regex, _ := regexp.Compile(`^\d{9}$`)
		return regex.MatchStrring(x)

	}, "cid": func(x string) bool {
		return true
	},
}

func main() {

	inputs, _ := reader("test.txt")

}
