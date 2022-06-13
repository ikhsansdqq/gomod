package main

import "fmt"

var i = 1

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Print(modRecursive(a, b))
}

func mod(x, y int) int {
	for x >= y {
		x -= y
	}

	return x
}

func modRecursive(a, b int) int {
	if b*i == a {
		return 0
	} else if b*i < a {
		return modRecursive(a, b)
	} else {
		return a - b*(i-1)
	}
}
