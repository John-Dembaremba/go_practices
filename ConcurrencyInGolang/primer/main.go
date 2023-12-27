package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {

	isPrime := func(randNum int) bool {

		for i := randNum - 1; i > 1; i-- {
			if randNum%i == 0 {
				return false
			}
		}
		return true
	}

	primeCh := make(chan int)

	go func() {
		defer close(primeCh)
		for {
			select {
			case <-done:
				return
			case randNum := <-randIntStream:
				if isPrime(randNum) {
					primeCh <- randNum
				}
			}
		}
	}()
	return primeCh
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	fannedStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for randNum := range c {
			select {
			case <-done:
				return
			case fannedStream <- randNum:
			}
		}
	}

	for _, primeCh := range channels {
		wg.Add(1)
		go transfer(primeCh)
	}

	go func() {
		defer close(fannedStream)
		wg.Wait()
	}()

	return fannedStream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 1; i <= n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)
	randNumGen := func() int { return rand.Intn(500000000) }

	// generator stage
	randNumStreamCh := repeatFunc(done, randNumGen)

	// naive approach
	// for randNum := range take(done, primeCh, 2) {
	// 	fmt.Println(randNum)
	// }

	// fan out stage
	//Assign gouritenes to number of CPUs
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randNumStreamCh)
	}

	// fan in stage
	primeCh := fanIn(done, primeFinderChannels...)

	// take stage
	for randNum := range take(done, primeCh, 10) {
		fmt.Println(randNum)
	}
	fmt.Println(time.Since(start))
}
