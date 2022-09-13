package workersPool

import (
	"sync"
)

var Inputs = make(chan Job, 10)
var Outputs = make(chan Result, 10)

type Job struct {
	Id string
	Inp int
}

type Result struct {
	Job Job
	Ans int
	WorkerNum int
}

func answer(num int) (ret int) {
	// 14
	for num != 0 {
		ret += (num%10)
		num /= 10;
	}
	return
}

func worker(wg *sync.WaitGroup, workerNum int) {
	for input := range Inputs {
		ans := answer(input.Inp)
		Outputs <- Result{input, ans, workerNum}
	}
	wg.Done()
}

func WorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i:=0; i<numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
}