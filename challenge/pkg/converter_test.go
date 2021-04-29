package pkg

import (
	"testing"
)

func TestConverterAsEnglishWillReturnEmtpyStringIfNumberIsLowerThanZero(t *testing.T) {
	result, ok := Converter(-1, AsEnglish)

	if ok == nil {
		t.Error("We expected an error since -1 is out of range")
	}

	if result != "" {
		t.Error("Return of function should be empty.")
	}
}

func TestConverterAsRomanWillReturnEmtpyStringIfNumberIsLowerThanOne(t *testing.T) {
	result, ok := Converter(0, AsRoman)

	if ok == nil {
		t.Error("We expected an error since 0 is out of range")
	}

	if result != "" {
		t.Error("Return of function should be empty.")
	}
}

func TestConverterAsEnglishWillReturnEmptyStringIfNumberIsGreaterThan(t *testing.T) {
	result, ok := Converter(99991, AsRoman)
	if ok == nil {
		t.Error("We expected an error since 99991 is out of range")
	}

	if result != "" {
		t.Error("Return of function should be empty.")
	}
}

func TestConvertToEnglishNumeralTests(t *testing.T) {
	testcases := []struct {
		number         int
		englishNumeral string
	}{
		{0, "Zero"},
		{1, "One"},
		{1993, "One thousand nine hundred and ninety-three"},
		{123, "One hundred and twenty-three"},
		{1234, "One thousand two hundred and thirty-four"},
		{85000, "Eighty-five thousand"},
	}

	for _, testcase := range testcases {
		result, ok := Converter(testcase.number, AsEnglish)
		if result != testcase.englishNumeral {
			t.Errorf("%d expected to convert to %s, got: %s.", testcase.number, testcase.englishNumeral, result)
		}
		if ok != nil {
			t.Errorf("%d convertion should not have an error. got: %s.", testcase.number, ok)
		}
	}
}

func TestConvertDebug(t *testing.T) {
	result, ok := Converter(1234, AsEnglish)

	if ok != nil {
		t.Error("Converting 1234 to string should not raise an error")
	}

	if result != "One thousand two hundred and thirty-four" {
		t.Errorf("We expected to see the string One thousand two hundred and thirty-four.Obtained \"%s\"", result)
	}
}

func TestConvertToRomanNumeralTest(t *testing.T) {
	testcases := []struct {
		number int
		roman  string
	}{
		{0, ""},
		{1, "I"},
		{2, "II"},
		{4, "IV"},
		{5, "V"},
		{1993, "MCMXCIII"},
		{2018, "MMXVIII"},
		{1111, "MCXI"},
		{2222, "MMCCXXII"},
		{444, "CDXLIV"},
		{555, "DLV"},
		{666, "DCLXVI"},
		{999, "CMXCIX"},
		{9999, "MMMMMMMMMCMXCIX"},
	}

	for _, testcase := range testcases {
		result, ok := Converter(testcase.number, AsRoman)
		if result != testcase.roman {
			t.Errorf("%d expected to convert to %s, got: %s.", testcase.number, testcase.roman, result)
		}
		if ok != nil {
			t.Errorf("%d convertion should not have an error. got: %s.", testcase.number, ok)
		}
	}
}
