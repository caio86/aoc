package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Numbers int

const (
	One Numbers = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

var numberMap = map[string]Numbers{
	"one":   One,
	"two":   Two,
	"three": Three,
	"four":  Four,
	"five":  Five,
	"six":   Six,
	"seven": Seven,
	"eight": Eight,
	"nine":  Nine,
}

func parseStringNumberToRealNumber(inputString string) string {
	var updatedString string = inputString
	var positions = make(map[int]string)

	for key := range numberMap {
		if strings.Contains(inputString, key) {
			positions[strings.Index(inputString, key)] = key
		}
	}

	positionKeys := make([]int, 0)
	for keys := range positions {
		positionKeys = append(positionKeys, keys)
	}
	sort.Ints(positionKeys)

	for _, key := range positionKeys {
		old := positions[key]
		replaceWith := fmt.Sprint(numberMap[positions[key]])
		updatedString = strings.Replace(updatedString, old, replaceWith, 1)
	}

	return updatedString
}

func getDigits(input string) string {
	var digits []rune

	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	return string(digits)
}

func calculateP1(inputScanner *bufio.Scanner) (sum uint) {
	var calibrationValues []uint

	for inputScanner.Scan() {
		input := inputScanner.Text()

		digits := getDigits(input)

		firstAndLastDigit, _ := strconv.ParseUint(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 64)

		calibrationValues = append(calibrationValues, uint(firstAndLastDigit))
	}

	for _, value := range calibrationValues {
		sum = sum + value
	}

	return
}

func calculateP2(inputScanner *bufio.Scanner) (sum uint) {
	var calibrationValues []uint

	for inputScanner.Scan() {
		input := inputScanner.Text()

		fmt.Printf("\ninput: %s\n", input)

		updatedInput := parseStringNumberToRealNumber(input)

		fmt.Printf("updatedInput: %s\n", updatedInput)

		digits := getDigits(updatedInput)

		fmt.Printf("digits: %v\n", digits)

		firstAndLastDigit, _ := strconv.ParseUint(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 64)

		fmt.Printf("first and last: %d\n", firstAndLastDigit)

		calibrationValues = append(calibrationValues, uint(firstAndLastDigit))
	}

	for _, value := range calibrationValues {
		sum = sum + value
	}

	return
}

func main() {
	inputFile := "./input.txt"
	file1, _ := os.Open(inputFile)
	defer file1.Close()

	file2, _ := os.Open(inputFile)
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	part1Result := calculateP1(scanner1)
	part2Result := calculateP2(scanner2)

	fmt.Printf("Part 1: %d\n", part1Result)
	fmt.Printf("Part 2: %d\n", part2Result)
	fmt.Printf("Test: %v\n", parseStringNumberToRealNumber("8kgplfhvtvqpfsblddnineoneighthg"))
	// Should result 88
}
