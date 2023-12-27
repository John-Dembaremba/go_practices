// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func runFanIn() {

// 	primeNumsCh := generateWork([]int{0, 3, 5, 7}, true)
// 	evenNumsCh := generateWork([]int{1, 2, 4, 6}, false)

// 	nums := fanIn(primeNumsCh, evenNumsCh)

// 	for num := range nums {
// 		fmt.Printf("Value: %v\n", num)
// 	}

// }

// func generateWork(nums []int, isPrime bool) <-chan int {
// 	ch := make(chan int)

// 	go func() {
// 		defer close(ch)

// 		for _, num := range nums {
// 			if isPrime {
// 				fmt.Printf("Writing %v number to PrimeNumbers channel.\n", num)
// 			} else {
// 				fmt.Printf("Writing %v number to EvenNumbers channel.\n", num)
// 			}
// 			ch <- num
// 		}
// 	}()
// 	fmt.Printf("Channel closed.\n")
// 	return ch
// }

// func fanIn(channels ...<-chan int) <-chan int {
// 	var wg sync.WaitGroup
// 	outPut := make(chan int)
// 	wg.Add(len(channels))

// 	for _, channel := range channels {
// 		go func(ch <-chan int) {
// 			for {
// 				num, isOpen := <-ch
// 				if !isOpen {
// 					fmt.Printf("Child Channel closed.\n")
// 					wg.Done()
// 					break
// 				}
// 				fmt.Printf("Writing %v number to Main channel.\n", num)
// 				outPut <- num
// 			}
// 		}(channel)
// 	}
// 	go func() {
// 		fmt.Printf("Closing Main channel.\n")
// 		wg.Wait()
// 		close(outPut)
// 	}()

// 	return outPut
// }
