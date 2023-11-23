package main

import "log"

func main() {

	log.Println(isPalindrome(121))

}

func isPalindrome(x int) bool {
	var reversed int
	temp := x

	for temp > 0 {
		reversed = reversed*10 + temp%10
		temp = temp / 10
	}
	return x == reversed
}
