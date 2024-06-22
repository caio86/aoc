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

func reverseString(input string) (result string) {
	for _, value := range input {
		result = string(value) + result
	}

	return
}

func parseStringNumberToRealNumber(inputString string) (result string, reversedResult string) {
	result = inputString
	reversedResult = inputString
	var positions = make(map[int]string)

	for index := 0; index <= 1; index++ {
		for key := range numberMap {
			if strings.Contains(inputString, key) {
				positions[strings.Index(inputString, key)] = key
			}
		}

		positionKeys := make([]int, 0)
		for keys := range positions {
			positionKeys = append(positionKeys, keys)
		}
		if index == 0 {
			sort.Ints(positionKeys)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(positionKeys)))
		}

		for _, key := range positionKeys {
			old := positions[key]
			replaceWith := fmt.Sprint(numberMap[positions[key]])
			if index == 0 {
				result = strings.Replace(result, old, replaceWith, 1)
			} else {
				old = reverseString(old)
				replaceWith = reverseString(replaceWith)
				reversedResult = reverseString(reversedResult)
				reversedResult = strings.Replace(reversedResult, old, replaceWith, 1)
				reversedResult = reverseString(reversedResult)
			}
		}
	}

	return
}

func getDigitsPart1(input string) string {
	var digits []rune

	// First digit
	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			break
		}
	}

	// Second digit
	for _, char := range reverseString(input) {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			break
		}
	}

	return string(digits)
}

func getDigitsPart2(input string) string {
	var digits []rune
	// input = "oooneeone"

	updatedInput, reversedInput := parseStringNumberToRealNumber(input)
	// First digit
	for _, char := range updatedInput {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			break
		}
	}

	// Second digit
	for _, char := range reverseString(reversedInput) {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			break
		}
	}

	return string(digits)
}

func calculateP1(inputScanner *bufio.Scanner) (sum uint) {
	var calibrationValues []uint

	for inputScanner.Scan() {
		input := inputScanner.Text()

		digits := getDigitsPart1(input)

		firstAndLastDigit, _ := strconv.ParseUint(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 64)

		calibrationValues = append(calibrationValues, uint(firstAndLastDigit))
	}

	for _, value := range calibrationValues {
		sum = sum + value
	}

	return
}

func calculateP2(inputScanner *bufio.Scanner) (sum uint64) {
	var calibrationValues []uint64

	for inputScanner.Scan() {
		input := inputScanner.Text()

		fmt.Printf("\ninput: %s\n", input)

		digits := getDigitsPart2(input)

		fmt.Printf("digits: %v\n", digits)

		firstDigit := digits[0]
		secondDigit := digits[len(digits)-1]

		var firstAndLastDigit uint64

		if firstDigit != secondDigit {
			firstAndLastDigit, _ = strconv.ParseUint(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 64)
		} else {
			firstAndLastDigit, _ = strconv.ParseUint(fmt.Sprintf("%c", digits[0]), 10, 64)
		}

		fmt.Printf("first and last: %d\n", firstAndLastDigit)

		calibrationValues = append(calibrationValues, firstAndLastDigit)
	}

	for _, value := range calibrationValues {
		sum = sum + value
	}

	return
}

func main() {
	inputFile := "../input.txt"
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
}
