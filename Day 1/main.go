package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	total, max := 0, 0
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := calc(mass)
		total += fuel

		for fm := calc(mass); fm > 0; {
			max += fm
			fm = calc(fm)
		}
	}

	fmt.Printf("%s%d\n", "Part 1: ", total)
	fmt.Printf("%s%d\n", "Part 2: ", max)
}

func calc(mass int) int {
	return mass/3 - 2
}
