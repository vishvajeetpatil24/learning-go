package other

import "fmt"

func OutputPerType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("The type is INT")
	case string:
		fmt.Println("The type is String")
	case []int:
		fmt.Println("The type is integer slice")
	}
}