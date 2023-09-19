package main

import "strings"

func intToRoman(num int) string {
	mapping := make(map[int]string)
	mapping[1] = "I"
	mapping[4] = "IV"
	mapping[5] = "V"
	mapping[9] = "IX"
	mapping[10] = "X"
	mapping[40] = "XL"
	mapping[50] = "L"
	mapping[90] = "XC"
	mapping[100] = "C"
	mapping[400] = "CD"
	mapping[500] = "D"
	mapping[900] = "CM"
	mapping[1000] = "M"

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for _, v := range nums {
		if num >= v {
			r := num / v
			result = result + strings.Repeat(mapping[v], r)
			num = num % v
		}
	}

	return result
}
