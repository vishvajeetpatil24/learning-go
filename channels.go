package main

import (
	"fmt"
	"time"
)

func sum(ch chan int, a ...int) {
	var ans int
	for _, v := range a {
		ans += v
	}
	time.Sleep(5 * time.Second)
	ch <- ans
	ch <- 5
	// We never reach this that means send and recieve both are blocking in case of unbuffered channels
	// But we reach here in case of buffered channels
	fmt.Println("Printing after sending two values to channels in sum function")
	ch <- 6
	fmt.Println("Printing after sending three values to channels in sum function")
	ch <- 7
	fmt.Println("Printing after sending four values to channels in sum function")
	close(ch)
}

func mult(ch chan int, a ...int) {
	var ans int = 1
	for _, v := range a {
		ans *= v
	}
	time.Sleep(5 * time.Second)
	ch <- ans
	ch <- 0
	fmt.Println("Printing after sending two valeus to channels in mult function")
	ch <- 6
	fmt.Println("Printing after sending three values to channels in mult function")
	ch <- 7
	fmt.Println("Printing after sending four values to channels in mult function")
	close(ch)
}

// Deadlock functions
func deadFirst(send chan int, receive chan int) {
	send <- 5
	<-receive
}

func deadSecond(send chan int, receive chan int) {
	<-receive
	send <- 8
}

// Deadlock functions
func deadThird(send chan int, receive chan int) {
	send <- 5
	<-receive
}

func deadFourth(send chan int, receive chan int) {
	send <- 8
	<-receive
}
