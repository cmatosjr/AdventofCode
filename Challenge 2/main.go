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

func main() {
	ranges, err := pars("./input.txt")
	if err != nil {
		panic(err)
	}
	for i := range ranges {
		for ranges[i][0] < ranges[i][1] {

		}
	}
}
