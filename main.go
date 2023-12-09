package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	day1()
}

func day1() {
	f, err := os.Open("./ChallengeFiles/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	calibrationSum := 0

	for scanner.Scan() {
		// parse line and put each number into an array

		var calibrationDigits []int
		line := scanner.Text()

		for _, char := range line {

			num, charErr := CharToNum(char)

			if charErr != nil {
				// do nothing
			} else {
				// append number to calibrationDigits
				calibrationDigits = append(calibrationDigits, num)
			}

		}

		if len(calibrationDigits) > 0 {
			var calibrationValue int

			// first digit is calibrationDigits[0]
			firstDigit := calibrationDigits[0] * 10

			// second digit is calibrationDigits[-1]
			lastDigit := calibrationDigits[len(calibrationDigits)-1] // does empty array break here?

			calibrationValue = firstDigit + lastDigit

			calibrationSum += calibrationValue
		}
		fmt.Println(calibrationSum)
	}
}

var ErrRuneNotInt = errors.New("type: rune was not int")

func CharToNum(r rune) (int, error) {
	if '0' <= r && r <= '9' {
		return int(r) - '0', nil
	}
	return 0, ErrRuneNotInt
}
