package soundex

import "strings"

var toIgnore = []string{"A", "E", "I", "O", "U", "H", "W", "Y"}
var one = []string{"B", "F", "P", "V"}
var two = []string{"C", "G", "J", "K", "Q", "S", "X", "Z"}
var three = []string{"D", "T"}
var four = []string{"L"}
var five = []string{"M", "N"}
var six = []string{"R"}

func Soundex(input string) string {
	upperCaseInput := strings.ToUpper(input)

	//split first letter
	firstLetter := string(upperCaseInput[0])
	numberPart := strings.Replace(upperCaseInput, firstLetter, "", 1)

	//soundex string replacements
	numberPart = replaceAllArrayElements(numberPart, toIgnore, "")
	numberPart = replaceAllArrayElements(numberPart, one, "1")
	numberPart = replaceAllArrayElements(numberPart, two, "2")
	numberPart = replaceAllArrayElements(numberPart, three, "3")
	numberPart = replaceAllArrayElements(numberPart, four, "4")
	numberPart = replaceAllArrayElements(numberPart, five, "5")
	numberPart = replaceAllArrayElements(numberPart, six, "6")

	//remove duplicates
	numberPart = removeDuplicates(numberPart)

	//add zeros if result string too short
	numbersOfZerosToAdd := 3 - len(numberPart)
	for index := 0; index < numbersOfZerosToAdd; index++ {
		numberPart += "0"
	}

	return firstLetter + numberPart
}

func replaceAllArrayElements(input string, array []string, replacement string) string {
	for _, letter := range array {
		input = strings.ReplaceAll(input, letter, replacement)
	}
	return input
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
