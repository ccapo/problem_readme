package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compute(threshold float64, limit float64) {
	var vsum float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return
		}

		delta := min(limit-vsum, max(0.0, value-threshold))
		vsum += delta

		fmt.Printf("%.1f\n", delta)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}

	fmt.Printf("%.1f\n", vsum)
}

func main() {
	// Check if we have the correct number of command-line arguments
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], " threshold limit")
		return
	}

	// Parse the command-line arguments
	threshold_str := os.Args[1]
	threshold, err := strconv.ParseFloat(threshold_str, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	limit_str := os.Args[2]
	limit, err := strconv.ParseFloat(limit_str, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}

	// Compute the sum
	compute(threshold, limit)
}
