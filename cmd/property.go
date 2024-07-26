package main

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var romanNumerals = []RomanNumeral{
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

func ConvertToRoman(val uint16) string {
	if val > 3999 {
		return "OoOoOoOoOoOo"
	}
	var result strings.Builder

	for _, romanNumeral := range romanNumerals {
		for val >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			val -= romanNumeral.Value
		}
	}
	return result.String()
}

func ConvertToArabicNumeral(romanNumeral string) uint16 {
	var arabicNumeral uint16 = 0

	for _, romeSymbol := range romanNumerals {
		for strings.HasPrefix(romanNumeral, romeSymbol.Symbol) {
			arabicNumeral += romeSymbol.Value
			romanNumeral = strings.TrimPrefix(romanNumeral, romeSymbol.Symbol)
		}
	}

	return arabicNumeral
}
