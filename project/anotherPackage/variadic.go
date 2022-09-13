package anotherPackage

import "fmt"

func ITakeMany(a int, b string, c, d int, e ...string) []string{
	fmt.Println("String b is", b)
	fmt.Println("The sum is", a + c + d)
	fmt.Printf("The type of e is %T\n", e)
	return e
}