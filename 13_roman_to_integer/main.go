package main

import "log"

func main() {

	log.Println(romanToInt("MX"))
}

func romanToInt(s string) int {

	charMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return charMap[s[0]]
	}

	sum := charMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if charMap[s[i]] < charMap[s[i+1]] {
			sum = sum - charMap[s[i]]
		} else {
			sum = sum + charMap[s[i]]
		}
	}
	return sum
}
