package main

import . "fmt"

func main() {
	var (
		opt   string
		point int
	)

	Println("What type of operation you desire?\n 1. String\n 2. Numbers")
	Scanln(&opt)

	if opt == string("string") {
		Println("Please write your words: ")
		Scanln(&opt)
		Println(isPalindrome(opt))
		Println(reverse(opt))
	} else if opt == string("numbers") {
		Println("Please write your numbers: ")
		Scanln(&point)
		Println(isPalindrome(string(rune(point))))
	} else {
		Println("You did not choose the right option, operation abort.")
	}
}

func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func strPalindrome(st string) string {
	return st
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

//TODO: CREATE PALINDROME
