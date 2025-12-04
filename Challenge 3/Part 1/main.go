package main

import (
	"bufio"
	"fmt"
	"os"
)

func pars(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := pars("input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for i, line := range lines {
		high_num := 0

		for j := 0; j < len(line); j++ {
			ch1 := line[j]
			num1 := int(ch1-'0') * 10

			for k := j + 1; k < len(line); k++ {
				ch2 := line[k]
				num2 := int(ch2 - '0')

				num3 := num1 + num2
				if num3 > high_num {
					high_num = num3
				}
			}
		}

		fmt.Printf("Line %d: %d\n", i, high_num)
		total += high_num
	}
	fmt.Printf("Total: %d\n", total)
}
