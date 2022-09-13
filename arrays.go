package main

var FirstArr = [5]int{1,2,3,4,5}
var prime [1000000]int

func EulerSieve() {
	var marker [1000000]int
	for i:= 2; i<=1000; i++ {
		if marker[i] == 1 {
			continue
		}
		for j:= 2; i*j<1000000; j++ {
			marker[i*j] = 1;
		}
	}
	index := 0
	for i:= 2; i<1000000; i++ {
		if marker[i] == 0 {
			prime[index] = i
			index++
		}
	}
}

func ChangeMiddle(arr [3]int) [3]int{
	len := len(arr)
	len = len/2
	arr[len] = 0
	return arr
}