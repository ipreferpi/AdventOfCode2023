package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//day1Solution := day1.Solve()
	//fmt.Printf("Day 1 Solution:\n%v\n", day1Solution)

	day2Solution := day2Solve()
	fmt.Printf("Day 2 Solution:\n%v\n", day2Solution)
}

func day2Solve() int {
	// open file
	f, err := os.Open("./challenges/day2/resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// new scanner to read lines individually
	scanner := bufio.NewScanner(f)

	possibleGamesIDSum := 0
	currentID := 1

	for scanner.Scan() {
		possible := true
		line := scanner.Text()

		// split string after :
		game := strings.Split(line, ":")[1]

		pulls := strings.Split(game, ";")
		//fmt.Println(pulls)

		for _, draw := range pulls {
			kek := strings.Split(draw, ",")
			kek = Map(kek, func(item string) string { return strings.Replace(item, " ", "", -1) })

			for _, fuck := range kek {
				if strings.Contains(fuck, "red") {
					numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "red", ""))
					if numGrabbed > 12 {
						possible = false
					}
				} else if strings.Contains(fuck, "green") {
					numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "green", ""))
					if numGrabbed > 13 {
						possible = false
					}
				} else if strings.Contains(fuck, "blue") {
					numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "blue", ""))
					if numGrabbed > 14 {
						possible = false
					}
				}
			}

		}

		fmt.Println(pulls)

		// split into each game, split at ;

		if possible {
			possibleGamesIDSum += currentID
		}

		fmt.Printf("Game #%v possibility is %v\n", currentID, possible)
		fmt.Println(possibleGamesIDSum)
		currentID++

	}

	return possibleGamesIDSum
}

// Generic Map function pulled here: https://stackoverflow.com/a/72498530

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func checkPossibleGame(line string) {

}
