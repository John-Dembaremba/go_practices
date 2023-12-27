package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fanOut(_ctx context.Context, wg *sync.WaitGroup, resultCh chan interface{}, primeNums, evenNums, oldNums int, primeCh <-chan int, evenCh <-chan int, oldCh <-chan int) {
	_, cancel := context.WithCancel(_ctx)
	defer cancel()
	defer wg.Done()

	wgPrime := &sync.WaitGroup{}
	wgEven := &sync.WaitGroup{}
	wgOld := &sync.WaitGroup{}

	wgPrime.Add(1)
	go func() {
		defer wgPrime.Done()
		data := make(map[string]int)
		for i := 0; i < primeNums; i++ {
			key := fmt.Sprintf("%d| Prime", i)
			data[key] = <-primeCh
			resultCh <- data
		}
	}()

	wgEven.Add(1)
	go func() {
		defer wgEven.Done()
		data := make(map[string]int)
		for i := 0; i < evenNums; i++ {
			key := fmt.Sprintf("%d| Even", i)
			data[key] = <-evenCh
			resultCh <- data
		}
	}()

	wgOld.Add(1)
	go func() {
		defer wgOld.Done()
		data := make(map[string]int)
		for i := 0; i < oldNums; i++ {
			key := fmt.Sprintf("%d| Old", i)
			data[key] = <-oldCh
			resultCh <- data
		}
	}()

	// Wait for all goroutines to finish concurrently
	wgPrime.Wait()
	wgEven.Wait()
	wgOld.Wait()
}

func fanIn(ctx context.Context, mainCh <-chan int, wg *sync.WaitGroup) (<-chan int, <-chan int, <-chan int) {
	primeCh, evenCh, oldCh := make(chan int), make(chan int), make(chan int)
	wgCategories := &sync.WaitGroup{}

	// Start goroutines for each category
	startGoroutine := func(ch chan int, category string) {
		defer wgCategories.Done()
		defer close(ch) // Close the channel when done
		for {
			select {
			case <-ctx.Done():
				return
			case num, ok := <-mainCh:
				if !ok {
					return
				}
				if isPrime(num) && category == "prime" {
					ch <- num
				} else if isDivisibleByTwo(num) && category == "even" {
					ch <- num
				} else if !isDivisibleByTwo(num) && category == "old" {
					ch <- num
				}
			}
		}
	}

	// Start goroutines for each category
	wgCategories.Add(3)
	go startGoroutine(primeCh, "prime")
	go startGoroutine(evenCh, "even")
	go startGoroutine(oldCh, "old")

	// Wait for all goroutines to finish concurrently
	go func() {
		wgCategories.Wait()
		close(primeCh)
		close(evenCh)
		close(oldCh)
	}()

	return primeCh, evenCh, oldCh
}

func concurrent(randNums, primeNums, evenNums, oldNums int) {
	var wg sync.WaitGroup
	ctx := context.Background()
	mainCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mainCh <- generateRandomNum(randNums)
			}
		}
	}()

	primeCh, evenCh, oldCh := fanIn(ctx, mainCh, &wg)

	resultCh := make(chan interface{})

	fanOut(ctx, &wg, resultCh, primeNums, evenNums, oldNums, primeCh, evenCh, oldCh)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data, ok := <-resultCh
				if !ok {
					return
				}

				fmt.Printf("Received data type: %T\n", data)

				if result, ok := data.(map[string]int); ok {
					for key, value := range result {
						fmt.Printf("Data %s ==>> %d\n", key, value)
					}
				} else {
					fmt.Println("Unexpected data type received")
				}
			}
		}
	}()
	wg.Wait()

}

func invoke() {

	randNums, primeNums, evenNums, oldNums := 1000000000, 3, 4, 4
	startTime := time.Now()
	isSequential := false
	if isSequential {
		sequential(randNums, primeNums, evenNums, oldNums)
	} else {
		concurrent(randNums, primeNums, evenNums, oldNums)
	}
	fmt.Println(time.Since(startTime))
}
