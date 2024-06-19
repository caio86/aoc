package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getDigits(input string) string {
	var digits []rune
	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	return string(digits)
}

func calculatePart1(inputScanner *bufio.Scanner) (sum uint) {
	var calibrationValues []uint
	for inputScanner.Scan() {
		digits := getDigits(inputScanner.Text())

		firstAndLastDigit, _ := strconv.ParseUint(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 64)

		calibrationValues = append(calibrationValues, uint(firstAndLastDigit))
	}

	for _, value := range calibrationValues {
		sum = sum + value
	}

	return
}

func main() {
	file, err := os.Open("./input_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Result := calculatePart1(scanner)

	fmt.Printf("Part 1 result: %d\n", part1Result)
}
