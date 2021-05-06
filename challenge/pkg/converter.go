package pkg

import (
	"errors"
	"log"
	"math"
	"unicode"
)

const (
	AsEnglish          rune = 'e'
	AsRoman            rune = 'r'
	minimumNumber      int  = 0
	maximumNumber      int  = 999999999
	minimumRomanNumber int  = 1
	maximumRomanNumber int  = 9999
)

// how many digit's groups to process
const groupsNumber int = 3

// getSmallsNumbers Return an array of strings.
func getSmallNumbers() []string {
	return []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}
}

func getTens() []string {
	return []string{
		"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety",
	}
}

func getScaleNumbers() []string {
	return []string{"", "thousand", "million"}
}

// Converter will take an int and will provide the value the english numeral as a string.
// If there is an error, the output tuple will contain an error.
// Since signature of func is int, octal will be converted automatically. 0123 8 == 83 10
// number is the number that you want to see represented as string representation
// concatenateWithAnd joins numbers with the string and
func Converter(number int, selector rune) (string, error) {
	switch selector {
	case AsRoman:
		if number < minimumRomanNumber || number > maximumRomanNumber {
			log.Printf("[pkg.converter.Converter] %d is not a valid input for conversor type %c", number,
				selector)
			return "", errors.New("number out of range")
		}
		return convertToRoman(number)
	case AsEnglish:
		if number < minimumNumber || number > maximumNumber {
			log.Printf("[pkg.converter.Converter] %d is not a valid input.", number)
			return "", errors.New("number out of range")
		}
		return convertToEnglishNumeral(number), nil
	}

	return "", errors.New("unknown conversor")
}

// convertToRoman convert a number to roman
func convertToRoman(number int) (string, error) {
	var asRoman string
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, conversion := range conversions {
		for number >= conversion.value {
			asRoman += conversion.digit
			number -= conversion.value
		}
	}
	return asRoman, nil
}

// convertToEnglishNumeral Return the english numeral representation of a number as a string
func convertToEnglishNumeral(number int) string {
	useAnd := true
	// Zero rule
	if number == 0 {
		return capitalize(getSmallNumbers()[0])
	}

	// Divide into three-digits group
	var groups [groupsNumber]int
	positive := math.Abs(float64(number))

	// Form four-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = int(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroupToText(groups[i], useAnd)
	}
	combined := textGroup[0]
	and := groups[0] > 0 && groups[0] < 100

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + getScaleNumbers()[i]

			if len(combined) != 0 {
				prefix += separator(and)
			}

			and = false

			combined = prefix + combined
		}
	}

	return capitalize(combined)
}

// intMod Obtain the rest of the division as an integer
func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

// digitGroupToText Converts a group of integers to the string representation
func digitGroupToText(group int, useAnd bool) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(group, 100)

	if hundreds != 0 {
		ret += getSmallNumbers()[hundreds] + " hundred"

		if tensUnits != 0 {
			ret += separator(useAnd)
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += getTens()[tens]

		if units != 0 {
			ret += "-" + getSmallNumbers()[units]
		}
	} else if tensUnits != 0 {
		ret += getSmallNumbers()[tensUnits]
	}

	return
}

// separator returns proper separator string between
// number groups.
func separator(useAnd bool) string {
	if useAnd {
		return " and "
	}
	return " "
}

// capitalize Python implementation of capitalize
func capitalize(s string) string {
	for index, value := range s {
		return string(unicode.ToUpper(value)) + s[index+1:]
	}
	return ""
}
