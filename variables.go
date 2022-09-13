package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"bytedetective.net/workers/workersPool"
	"github.com/google/uuid"
	"github.com/myorg/project/anotherPackage"
	"github.com/myorg/project/other"

	"rsc.io/quote"
)

// type S struct {
// 	a, b int
// }

// // String implements the fmt.Stringer interface
// func (s *S) String() string {
// 	return fmt.Sprintf("%s", s) // Sprintf will call s.String()
// }

func main() {
	var x, y int = 5, 6
	fmt.Println(x, " ", y)
	z := strconv.Itoa(2839)
	fmt.Printf("The number in z as string is %s and type of same is %T\n", z, z)
	var a, b = "Russa", true
	fmt.Printf("The type of a is %T and of b is %T\n", a, b)
	fmt.Println(other.OtherFunc(3, 4))
	fmt.Println(quote.Glass())
	fmt.Println(MULTILINE)
	fmt.Printf("The type of untyped constant Pi is %T\n", Pi)

	// Checking other edge cases
	var u int
	var v []int
	// The below line will give error because it is declaration. And we already have declared u.
	// u := 10
	u = 10
	u = 43
	w, x := 22, 34
	l, x := 55, 14
	fmt.Println(u, " ", v, " ", w, " ", l, " ", x)

	// Operators (*, &, <-)
	// No pointer arithmatic in this language
	fmt.Printf("The address of u is %x\n", &u)
	// TODO: Add more here on receive operator

	// Conditions
	fmt.Println(conditionalLogic(8))

	// Array
	fmt.Println(FirstArr)

	// Array and loop (tested and works)
	EulerSieve()
	fmt.Println(prime[48437])

	var chMiddle = [3]int{3, 4, 5}
	newArr := ChangeMiddle(chMiddle)
	fmt.Println(newArr)

	// Slices
	slc := []int{3, 4, 5}
	ChangeSlice(slc)
	fmt.Println(slc)

	// chMiddle is unchanged array
	fmt.Println(chMiddle)
	// We will now create slice from it and will print again
	ChangeSlice(chMiddle[:])
	// Printing it again. Proves that any change in slice is reflected in its referenced array
	fmt.Println(chMiddle)

	// Slice from slice
	newSlc := slc[0:2]
	fmt.Println(newSlc)
	fmt.Println(cap(newSlc))
	fmt.Println(len(newSlc))
	newSlc = newSlc[0:3]
	fmt.Println(newSlc)
	fmt.Println(len(newSlc))

	// MultiDimensional arrays and slices
	newArr1 := [3][4]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}
	newArr2 := [3][3]string{{"India", "China", "Russia"}, {"A", "B", "C"}, {"X", "Y", "Z"}}
	fmt.Println(newArr1)
	fmt.Println(newArr2)

	newSlc1 := [][]int{{1, 2, 3}, {3, 4}, {6, 7, 8}}
	fmt.Println(newSlc1[0][1])
	fmt.Println(newSlc1[0:2][0:1])
	var nilSlc []int
	var nonNilSlc = []int{4, 5}
	fmt.Println(nilSlc == nil, " ", nonNilSlc == nil)

	// Variadic functions
	fmt.Println(anotherPackage.ITakeMany(5, "India", 2, 3, "Russian", "People", "are"))

	slcForVariadic := []string{"X", "Y", "Z"}
	fmt.Println(anotherPackage.ITakeMany(10, "China", 4, 5, slcForVariadic...))

	// Using anon functions
	fmt.Println(anotherPackage.UseAnon("The sum is", 5, 5))

	fmt.Println(anotherPackage.TakeAnon(anotherPackage.AnonMultiply, 4, 5))

	fmt.Println(anotherPackage.Singleton(3, 4, 5))

	fmt.Println(anotherPackage.Singleton(3, 4, 6))

	// Declaring and defining structures
	var newStructure anotherPackage.Structure = anotherPackage.Structure{FieldA: 5}
	var newImplicitStructure anotherPackage.Inter = &anotherPackage.ImplicitStructures{Structure: newStructure}

	newImplicitStructureDynamicVal := newImplicitStructure.(*(anotherPackage.ImplicitStructures))
	// Modifying newImplicitStructure from inside
	newImplicitStructure.Fourth()

	fmt.Println("NewImplicitStructure value and its int and string print")
	fmt.Println(*newImplicitStructureDynamicVal)
	fmt.Println(newImplicitStructureDynamicVal.Print())

	fmt.Println("NewStructure FieldA")
	fmt.Println(newStructure.FieldA)
	(&newStructure).Change(6)
	newStructure.Change(9)
	fmt.Println("NewStructure again after change")
	fmt.Println(newStructure)

	fmt.Println(*newImplicitStructureDynamicVal)
	fmt.Println(newImplicitStructureDynamicVal.Print())

	var interfaceDynamic anotherPackage.SampleInterface = anotherPackage.ImplicitStructures{}

	newInterfaceDynamicVal := interfaceDynamic.(anotherPackage.ImplicitStructures)

	newInterfaceDynamicVal.Fourth()
	fmt.Println(newInterfaceDynamicVal.Print())

	var xyz []int
	if xyz == nil {
		fmt.Println("xyz is nil")
	}
	other.OutputPerType(xyz)

	var slcWithoutMake []int = []int{4, 5}
	fmt.Println(slcWithoutMake[:][:][1])

	employeeSalary := map[string]int{}
	employeeSalary["Russia"] = 19
	employeeSalary["China"] = 10
	for key, value := range employeeSalary {
		fmt.Println(key, " ", value)
	}
	ch1, ch2 := make(chan int), make(chan int)
	go mult(ch1, 3, 4, 5)
	go sum(ch2, 3, 4, 5)
	fmt.Println(<-ch1 + <-ch2)
	time.Sleep(6 * time.Second)

	// Buffered channels
	ch3, ch4 := make(chan int, 2), make(chan int, 2)
	go mult(ch3, 1, 2, 3)
	go sum(ch4, 1, 2, 3)
	fmt.Println(<-ch3 + <-ch4)
	time.Sleep(1 * time.Second)

	// Closing channels and foreach loop with closed channels
	for value := range ch3 {
		fmt.Println("Sent values in channel ch3 is ", value)
	}

	// OS deadlock analogy (With unbuffered ones)
	ch5 := make(chan int)
	ch6 := make(chan int)
	// s := &S{a: 1, b: 2}
	// fmt.Println(s)

	// First value will be sent to ch5 and it will block on someone to read it but it is supposed
	// to be read from deadSecond which will then read it and then value will be blocked to be read
	// from ch6 but it will be then sent from deadSecond so this IS NOT DEADLOCK
	go deadFirst(ch5, ch6)
	go deadSecond(ch6, ch5)

	// There is order of operations reversed with deadFourth compared to deadSecond and hence
	// deadlock will end up happening if main subroutine is also waiting on something. Because for
	// deadlock to happen all subroutines are sleeping.
	ch7 := make(chan int)
	ch8 := make(chan int)
	// ch9 := make(chan int)
	go deadThird(ch7, ch8)
	go deadFourth(ch8, ch7)
	time.Sleep(1 * time.Second)

	_, err := SomeFunc(400)
	if err != nil {
		fmt.Println(err.Code)
	}
	go workersPool.WorkerPool(10)
	go (func() {
		for {
			for i := 0; i < 100; i++ {
				byteUUID := uuid.New()
				workersPool.Inputs <- workersPool.Job(workersPool.Job{Id: byteUUID.String(), Inp: rand.Int()})
			}
			time.Sleep(10 * time.Second)
		}
	})()

	for res := range workersPool.Outputs {
		fmt.Printf("The job id was %s and answer was %d for input %d and done by worker %d\n", res.Job.Id, res.Ans, res.Job.Inp, res.WorkerNum)
	}
	fmt.Println("Hey how come is this true")
}
