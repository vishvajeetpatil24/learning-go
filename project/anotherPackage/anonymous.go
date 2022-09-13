package anotherPackage

import "strconv"

var AnonSum = func(a, b int) int {
	return a+b
}

var AnonMultiply = func(a, b int) int {
	return a*b
}

func UseAnon(str string, a, b int) string {
	return str + " " + strconv.Itoa(AnonSum(a, b))
}

func TakeAnon(anon func(a, b int) int, x, y int) int {
	return anon(x, y)
}

// No OOPS in this language. Just to demonstrate closure concept.
var Singleton = (func () func (numbers...int) []int {
	var closureVar []int = nil

	return func (numbers...int) []int {
		if (closureVar == nil) {
			closureVar = numbers
		}
		return closureVar
	}
})()

