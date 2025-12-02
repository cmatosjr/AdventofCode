package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./combo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.Create("debug.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)

	scanner := bufio.NewScanner(file)
	starting_place := 50
	ending_place := 0
	count := 0
	warpcount := 0
	fmt.Fprintln(mw, "The dial starts by pointing at 50.")
	for scanner.Scan() {
		text := scanner.Text()
		rest, err := strconv.Atoi(text[1:])
		if err != nil {
			panic(err)
		}
		warpcount = warpcount + (rest / 100)
		result := rest % 100

		switch text[0] {
		case 'R':
			ending_place = starting_place + result
			if ending_place > 100 {
				warpcount++
			}
		case 'L':
			ending_place = starting_place - result
			if ending_place < 0 && starting_place != 0 {
				warpcount++
			}
		}
		ending_place = ((ending_place % 100) + 100) % 100

		if ending_place == 0 {
			count++
			warpcount++
		}
		fmt.Fprintf(mw, "The dial is rotated %s to point at %d. The warp counter is %d.\n",
			text, ending_place, warpcount)
		starting_place = ending_place
	}
	fmt.Fprintf(mw, "Final count: %d, warpcount: %d\n", count, warpcount)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
}
