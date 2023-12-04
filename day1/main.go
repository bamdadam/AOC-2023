package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("error while reading file: ", err)
	}
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		in := scanner.Text()
		inRev := ReverseString(in)
		re := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
		reRev := regexp.MustCompile("\\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
		num1 := re.FindString(in)
		num2 := reRev.FindString(inRev)
		var n1, n2 uint64
		n1, err := strconv.ParseUint(num1, 10, 32)
		if err != nil {
			n1, err = numericToInt(num1)
			if err != nil {
				fmt.Println("error while converting number: ", err)
			}
		}
		n2, err = strconv.ParseUint(num2, 10, 32)
		if err != nil {
			n2, err = numericToInt(num2)
			if err != nil {
				fmt.Println("error while converting number: ", err)
			}
		}
		res += int(n1)*10 + int(n2)
	}
	fmt.Println(res)
}

func numericToInt(match string) (uint64, error) {
	switch match {
	case "one", "eno":
		return 1, nil
	case "two", "owt":
		return 2, nil
	case "three", "eerht":
		return 3, nil
	case "four", "ruof":
		return 4, nil
	case "five", "evif":
		return 5, nil
	case "six", "xis":
		return 6, nil
	case "seven", "neves":
		return 7, nil
	case "eight", "thgie":
		return 8, nil
	case "nine", "enin":
		return 9, nil
	}
	return 0, fmt.Errorf("match not found: %s", match)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
