package pkg

import (
	"testing"
)

func TestConverterWillReturnEmtpyStringIfNumberIsLowerThanZero(t *testing.T) {
	concatenateWithAnd := true
	result, ok := Converter(-1, concatenateWithAnd)

	if ok == nil {
		t.Error("We expected an error since -1 is out of range")
	}

	if result != "" {
		t.Error("Return of function should be empty.")
	}
}

func TestConverterWillReturnEmptyStringIfNumberIsGreaterThan(t *testing.T) {
	concatenateWithAnd := true
	result, ok := Converter(9999999991, concatenateWithAnd)
	if ok == nil {
		t.Error("We expected an error since 9999999991 is out of range")
	}

	if result != "" {
		t.Error("Return of function should be empty.")
	}
}

func TestConvertTenToString(t *testing.T) {
	concatenateWithAnd := true
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
		result, ok := Converter(testcase.number, concatenateWithAnd)
		if result != testcase.englishNumeral {
			t.Errorf("%d expected to convert to %s, got: %s.", testcase.number, testcase.englishNumeral, result)
		}
		if ok != nil {
			t.Errorf("%d convertion should not have an error. got: %s.", testcase.number, ok)
		}
	}
}

func TestConvertDebug(t *testing.T) {
	concatenateWithAnd := true
	result, ok := Converter(1234, concatenateWithAnd)

	if ok != nil {
		t.Error("Converting 1234 to string should not raise an error")
	}

	if result != "One thousand two hundred and thirty-four" {
		t.Errorf("We expected to see the string One thousand two hundred and thirty-four.Obtained \"%s\"", result)
	}
}
