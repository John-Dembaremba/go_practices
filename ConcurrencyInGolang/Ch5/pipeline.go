package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNum(randNums int) int {
	return rand.Intn(randNums)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := num - 1; i >= 2; i-- {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isDivisibleByTwo(num int) bool {
	return num%2 == 0
}

func arrayNums(randNums, primeNums, evenNums, oldNums int) ([]int, []int, []int) {
	randNum := generateRandomNum(randNums)

	var primeSlice, evenSlice, oldSlice []int

	for {
		if isPrime(randNum) && len(primeSlice) < primeNums {
			primeSlice = append(primeSlice, randNum)
		} else if isDivisibleByTwo(randNum) && len(evenSlice) < evenNums {
			evenSlice = append(evenSlice, randNum)
		} else if !isDivisibleByTwo(randNum) && len(oldSlice) < oldNums {
			oldSlice = append(oldSlice, randNum)
		}
		if len(primeSlice) == primeNums && len(evenSlice) == evenNums && len(oldSlice) == oldNums {
			break
		}
	}
	return primeSlice, evenSlice, oldSlice

}

func sequential(randNums, primeNums, evenNums, oldNums int) {
	primeSlice, evenSlice, oldSlice := arrayNums(randNums, primeNums, evenNums, oldNums)

	fmt.Printf("Prime Numbers Slice: %v", primeSlice)
	fmt.Println()
	fmt.Printf("Even Numbers Slice: %v", evenSlice)
	fmt.Println()
	fmt.Printf("Prime Numbers Slice: %v", oldSlice)
	fmt.Println()
}
