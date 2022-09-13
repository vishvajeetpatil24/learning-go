package anotherPackage

import "fmt"

type Data int

type Structure struct {
	FieldA int
}

type ImplicitStructures struct {
	int
	string
	Structure
	FieldA int
}

func (str *Structure) Change(fieldA int) {
	(*str).FieldA = fieldA
}

func (str ImplicitStructures) Print() string {
	fmt.Println(str)
	ans := fmt.Sprintf("The given number is %d and given string is %s ", str.int, str.string)
	return ans
}

type Inter interface {
	first()
	second() int
	third(e...string) int
	Fourth()
}

func (str *ImplicitStructures) first() {
	(*str).int = 10
	(*str).string = "India"
}

func (str *ImplicitStructures) second() int {
	(*str).FieldA = 11
	return 11
}

func (str ImplicitStructures) third(e...string) int {
	fmt.Println(e)
	return -1
}

func (str *ImplicitStructures) Fourth() {
	str.first()
	str.second()
	str.third("Vishvajeet", "Subhash", "Patil")
}

type SampleInterface interface {
	a()
	b()
}

func (str ImplicitStructures) a() {
	fmt.Println("Method a")
}

func (str ImplicitStructures) b() {
	fmt.Println("Method b")
}