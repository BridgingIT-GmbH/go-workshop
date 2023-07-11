package soundex

import (
	"errors"
	"strings"
)

var soundexMapping = map[string]string{
	"B": "1",
	"F": "1",
	"P": "1",
	"V": "1",
	"C": "2",
	"G": "2",
	"J": "2",
	"K": "2",
	"Q": "2",
	"S": "2",
	"X": "2",
	"Z": "2",
	"D": "3",
	"T": "3",
	"L": "4",
	"M": "5",
	"N": "5",
	"R": "6",
}

const maxLengthNumberPart = 3

func Soundex(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input string")
	}
	upperCaseInput := strings.ToUpper(input)

	//split first letter
	firstLetter := string(upperCaseInput[0])
	remainingLetters := strings.Replace(upperCaseInput, firstLetter, "", 1)

	//soundex string replacements
	numberPart := ""
	for _, letter := range remainingLetters {
		numberPart += soundexMapping[string(letter)]
	}

	//remove duplicates
	numberPart = removeDuplicates(numberPart)

	//add zeros for short string and limit length to 3
	numberPart += "000"
	if len(numberPart) > maxLengthNumberPart {
		numberPart = numberPart[:maxLengthNumberPart]
	}

	return firstLetter + numberPart, nil
}

func removeDuplicates(input string) string {
	result := ""
	var lastLetter rune
	for _, letter := range input {
		if letter != lastLetter {
			result += string(letter)
			lastLetter = letter
		}
	}
	return result
}
