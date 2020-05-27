package main

import (
	"fmt"
	"strings"
)

var romanToNumerals = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func getNumerals(roman string) int {
	for r, n := range romanToNumerals {
		ss := strings.Split(roman, r)
		fmt.Println(ss, n)
	}

	return 0
}

func main() {
	getNumerals("XXI")
}
