package main

import (
	"adventofcode/challenges/day1"
	"adventofcode/challenges/day2"
	"fmt"
)

func main() {
	day1Solution := day1.Solve()
	fmt.Printf("Day 1 Part 1 Solution:\n%v\n", day1Solution)

	day2Part1, day2Part2 := day2.Solve()
	fmt.Printf("Day 2 Part 1 Solution:\n%v\n", day2Part1)
	fmt.Printf("Day 2 Part 2 Solution:\n%v\n", day2Part2)
}
