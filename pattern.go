package main

import "fmt"

type Pattern struct {
	match       string
	description string
}

func (p Pattern) String() string {
	return fmt.Sprintf("%s: %q", p.description, p.match)
}

type Patterns []Pattern

func (p Patterns) Strings() []string {
	ret := []string{}
	for _, pattern := range p {
		ret = append(ret, pattern.String())
	}
	return ret
}

func NewPattern(match, description string) Pattern {
	return Pattern{
		match:       match,
		description: description,
	}
}
