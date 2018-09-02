package main

import "strings"

type Report struct {
	input string
}

func Analyze(word string) Report {
	return Report{input: word}
}

func (r Report) Patterns() Patterns {
	patterns := Patterns{}

	// by groups
	for i := len(r.input) / 2; i > 1; i-- {
		parts := ChunkString(r.input[len(r.input)%i:], i)
		//fmt.Println(r.input, i, parts)

		// FIXME: do repetitions

		// palindrome
		// FIXME: support palindrome with odd amount of characters
		if len(parts) > 1 {
			if parts[len(parts)-1] == ReverseString(parts[len(parts)-2]) {
				patterns = append(patterns, NewPattern(
					strings.Join(parts[len(parts)-2:], ""),
					"palindrome",
				))
			}
		}

		// FIXME: find basic words in hexadecimal
		// FIXME: find famous numbers
	}

	return patterns
}

func (r Report) String() string {
	patterns := r.Patterns()
	if len(patterns) == 0 {
		return "no pattern found"
	}
	return strings.Join(patterns.Strings(), "; ")
}
