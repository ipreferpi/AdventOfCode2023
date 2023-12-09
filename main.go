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
	// open file
	f, err := os.Open("./ChallengeFiles/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// new scanner to read lines individually
	scanner := bufio.NewScanner(f)

	calibrationSum := 0

	for scanner.Scan() {

		calibrationValue := calcLineCalibrationValue(scanner.Text())
		calibrationSum += calibrationValue

	}
	fmt.Println(calibrationSum)
}

func calcLineCalibrationValue(line string) int {
	var calibrationDigits []int

	// iterate over each rune, and if it's a number, append it to calibrationDigits
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
		// first digit is calibrationDigits[0]
		firstDigit := calibrationDigits[0] * 10

		// second digit is calibrationDigits[-1]
		lastDigit := calibrationDigits[len(calibrationDigits)-1] // does empty array break here?

		return firstDigit + lastDigit
	}
	return 0
}

// https://codereview.stackexchange.com/a/122931
// nice solution to determining if a rune is a number, and throwing error if not.
// thanks peterSO!

var ErrRuneNotInt = errors.New("type: rune was not int")

func CharToNum(r rune) (int, error) {
	// This is just ascii number comparison, then subtracting 0's ascii value from the resulting int rune's ascii value
	if '0' <= r && r <= '9' {
		return int(r) - '0', nil
	}
	return 0, ErrRuneNotInt
}
