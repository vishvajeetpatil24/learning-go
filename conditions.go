package main

import "fmt"

func conditionalLogic(num int) bool {
	if num < 5 {
		return false
	}
	var ans bool = true
	if num % 2 == 0 {
		fmt.Println("Even Number")
		ans = true
	} else if num % 3 == 0 {
		fmt.Println("Divisible by 3")
		ans = false
	}
	return ans
}