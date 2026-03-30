package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func compute(writer io.Writer, threshold float64, limit float64, values []float64) {
	vsum := float64(0.0)

	// Iterate over input values
	for _, value := range values {
		// Transform values and update sum
		delta := min(limit-vsum, max(0.0, value-threshold))
		vsum += delta

		// Print each transformed value
		fmt.Fprintf(writer, "%.1f\n", delta)
	}

	// Print final sum
	fmt.Fprintf(writer, "%.1f\n", vsum)
}

func main() {
	// Check if we have the correct number of command-line arguments
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "threshold limit")
		return
	}

	// Parse threshold argument
	threshold_str := os.Args[1]
	threshold, err := strconv.ParseFloat(threshold_str, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	if threshold < 0.0 || threshold > 1000000000.0 {
	  fmt.Fprintln(os.Stderr, "error: threshold should be between 0.0 and 1000000000.0: ", threshold)
	  return
	}
	
	// Parse limit argument
	limit_str := os.Args[2]
	limit, err := strconv.ParseFloat(limit_str, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	if limit < 0.0 || limit > 1000000000.0 {
	  fmt.Fprintln(os.Stderr, "error: limit should be between 0.0 and 1000000000.0: ", limit)
	  return
	}

	// Parse values from STDIN
	values := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Convert string to float64
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return
		}

    // If in valid range, store input value
		if value >= 0.0 && value <= 1000000000.0 {
		  values = append(values, value)
		}
	}

	// Compute the sum
	compute(os.Stdout, threshold, limit, values)
}
