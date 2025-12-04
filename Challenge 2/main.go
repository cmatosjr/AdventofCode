package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pars(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("file is empty")
	}

	line := scanner.Text()
	parts := strings.Split(line, ",")

	var ranges [][]int

	for _, part := range parts {
		nums := strings.Split(strings.TrimSpace(part), "-")
		if len(nums) != 2 {
			return nil, fmt.Errorf("invalid range: %s", part)
		}

		start, err1 := strconv.Atoi(nums[0])
		end, err2 := strconv.Atoi(nums[1])

		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number in: %s", part)
		}

		ranges = append(ranges, []int{start, end})
	}

	return ranges, nil
}

func SplitDigits(n int) []int {
	s := strconv.Itoa(n)
	digits := make([]int, len(s))

	for i, ch := range s {
		digits[i] = int(ch - '0') // convert rune to int digit
	}

	return digits
}

/*
func hasRepeatingSequence(num string) bool {
	length := len(num)
	for size := 1; size <= length/2; size++ {
		seq := num[:size]
		if strings.Repeat(seq, length/size) == num {
			return true
		}
	}
	return false
}
*/

func isDoubleSequence(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	if length%2 != 0 {
		return false
	}

	half := length / 2
	first := s[:half]
	second := s[half:]

	return first == second
}

func main() {
	ranges, err := pars("./input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for i := range ranges {
		start := ranges[i][0]
		end := ranges[i][1]

		for num := start; num <= end; num++ {
			//str := strconv.Itoa(num)
			if isDoubleSequence(num) {
				total += num
				fmt.Printf("%d\n", num)
			}
		}
	}
	fmt.Printf("%d", total)
}
