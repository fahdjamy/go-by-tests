package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

const TestDescriptionTemplate = "given '%v' return '%v'"

var tableTestCases = []struct {
	Input uint16
	Want  string
}{
	{Input: 1, Want: "I"},
	{Input: 2, Want: "II"},
	{Input: 3, Want: "III"},
	{Input: 4, Want: "IV"},
	{Input: 5, Want: "V"},
	{Input: 9, Want: "IX"},
	{Input: 1006, Want: "MVI"},
	{Input: 1984, Want: "MCMLXXXIV"},
	{Input: 3999, Want: "MMMCMXCIX"},
	{Input: 798, Want: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	t.Run("covert number to roman numerals", func(t *testing.T) {
		for _, tt := range tableTestCases {
			t.Run(fmt.Sprintf(TestDescriptionTemplate, tt.Input, tt.Want), func(t *testing.T) {
				if got := ConvertToRoman(tt.Input); got != tt.Want {
					t.Errorf("got %q, want %q", got, tt.Want)
				}
			})
		}
	})

	t.Run("cover number to roman numerals", func(t *testing.T) {
		for _, tt := range tableTestCases {
			t.Run(fmt.Sprintf(TestDescriptionTemplate, tt.Want, tt.Input), func(t *testing.T) {
				if got := ConvertToArabicNumeral(tt.Want); got != tt.Input {
					t.Errorf("got %q, want %q", got, tt.Input)
				}
			})
		}
	})
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabicNumber uint16) bool {
		if arabicNumber < 0 || arabicNumber > 3999 {
			//log.Printf("arabic number out of range: %v", arabicNumber)
			return true
		}
		t.Log("testing", arabicNumber)
		roman := ConvertToRoman(arabicNumber)
		fromRoman := ConvertToArabicNumeral(roman)
		return fromRoman == arabicNumber
	}

	// The default number of runs quick.Check performs is 100, but you can change that with a config.
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("checks failed:", err)
	}
}
