package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	// open file
	f, err := os.Open("./challenges/day2/resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// new scanner to read lines individually
	scanner := bufio.NewScanner(f)

	possibleGamesIDSum := 0
	powerSum := 0
	currentID := 1

	for scanner.Scan() {
		line := scanner.Text()

		if checkPossibleGame(line) {
			possibleGamesIDSum += currentID
		}

		powerSum += checkGamePower(line)

		currentID++
	}
	return possibleGamesIDSum, powerSum
}

func checkPossibleGame(line string) bool {
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
					return false
				}
			} else if strings.Contains(fuck, "green") {
				numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "green", ""))
				if numGrabbed > 13 {
					return false
				}
			} else if strings.Contains(fuck, "blue") {
				numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "blue", ""))
				if numGrabbed > 14 {
					return false
				}
			}
		}
	}
	return true
}

func checkGamePower(line string) int {
	// split string after :
	game := strings.Split(line, ":")[1]

	pulls := strings.Split(game, ";")
	//fmt.Println(pulls)

	minRed := 0
	minGreen := 0
	minBlue := 0

	for _, draw := range pulls {
		kek := strings.Split(draw, ",")
		kek = Map(kek, func(item string) string { return strings.Replace(item, " ", "", -1) })

		for _, fuck := range kek {
			if strings.Contains(fuck, "red") {
				numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "red", ""))
				if numGrabbed > minRed {
					minRed = numGrabbed
				}
			} else if strings.Contains(fuck, "green") {
				numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "green", ""))
				if numGrabbed > minGreen {
					minGreen = numGrabbed
				}
			} else if strings.Contains(fuck, "blue") {
				numGrabbed, _ := strconv.Atoi(strings.ReplaceAll(fuck, "blue", ""))
				if numGrabbed > minBlue {
					minBlue = numGrabbed
				}
			}
		}
	}

	return minRed * minGreen * minBlue
}

// Generic Map function pulled here: https://stackoverflow.com/a/72498530

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
