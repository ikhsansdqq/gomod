package main

import "fmt"

func main() {
	var (
		opt   string
		point int
	)

	fmt.Println("What type of operation you desire?\n 1. String\n 2. Numbers")
	fmt.Print("Input: ")
	fmt.Scanln(&opt)

	if opt == string("string") {
		fmt.Println("Please write your words: ")
		fmt.Print("Input: ")
		fmt.Scanln(&opt)
		fmt.Println(isPalindrome(opt))
		fmt.Println(reverseString(opt))
	} else if opt == string("numbers") {
		fmt.Println("Please write your numbers: ")
		fmt.Print("Input: ")
		fmt.Scanln(&point)
		fmt.Println(isPalindrome(string(rune(point))))
		fmt.Println(reverseNumber(point))
	} else {
		fmt.Println("You did not choose the right option, operation abort.")
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

func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}

	return res
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
