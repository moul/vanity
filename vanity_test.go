package main

import "fmt"

func Example() {
	inputs := []string{
		"3A6666A3",
		"691ADDED",
		"72727212",
		"B42B02B4",
		"18473729",
		"00000000",
		"00000001",
		"10000000",
		"00100100",
		"10001000",
	}
	for _, input := range inputs {
		fmt.Println(input, Analyze(input))
	}

	// Output:
	// test
}
