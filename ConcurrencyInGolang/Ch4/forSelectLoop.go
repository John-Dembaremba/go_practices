// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func runner() {
// 	var wg sync.WaitGroup

// 	c1 := make(chan interface{}, 5)
// 	c2 := make(chan interface{}, 5)

// 	wg.Add(1)
// 	go func() {
// 		nums := []int{1, 2, 3, 4, 5}
// 		defer wg.Done()
// 		defer close(c1)

// 		for _, num := range nums {
// 			fmt.Printf("Pushing num: %v to %v channel ...\n", num, c1)
// 			c1 <- num
// 			time.Sleep(2 * time.Second)
// 			fmt.Printf("Pushing num: %v to %v channel completed.\n", num, c1)
// 		}
// 		fmt.Printf("%v channel closing", c1)
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		chars := []int{'a', 'b', 'c', 'd', 'e'}
// 		defer wg.Done()
// 		defer close(c2)

// 		for _, char := range chars {
// 			fmt.Printf("Pushing char: %v to %v channel ...\n", char, c2)
// 			c2 <- char
// 			time.Sleep(2 * time.Second)
// 			fmt.Printf("Pushing char: %v to %v channel completed.\n", char, c2)
// 		}
// 		fmt.Printf("%v channel closing", c2)
// 	}()

// 	for i := 0; i < 100; i++ {
// 		select {
// 		case num := <-c1:
// 			fmt.Printf("%v channel %v received\n", c1, num)

// 		case char := <-c2:
// 			fmt.Printf("%v channel %v received\n", c2, char)
// 		default:
// 			fmt.Println("All channels closed")
// 			wg.Wait()

// 		}
// 	}
// }
