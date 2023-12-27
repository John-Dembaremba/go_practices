package main

import (
	"fmt"
	"strconv"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Print(myArray[3])
	fmt.Println()
	fmt.Print(len(myArray))
	fmt.Println()
	fmt.Print(twoDimensionalArray())
	fmt.Println()
	fmt.Print(threeDimensionalArray())

	mySlice := make([]int, 5)
	s := []int{1, 2, 3}
	fmt.Println()
	fmt.Print(mySlice)
	fmt.Println()
	fmt.Print(s)
	fmt.Println()
	fmt.Println()
	fmt.Print(arraySlicing())
	fmt.Println()
	fmt.Println()
	fmt.Print(copyArrays())
	fmt.Println()
	fmt.Println()
	fmt.Print(copySlices())
}

func twoDimensionalArray() [5][6]string {
	var array [5][6]string

	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			array[row][column] = strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}

	return array
}

func threeDimensionalArray() [2][3][3]string {
	var cube [2][3][3]string
	for row := 0; row < 2; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] = strconv.Itoa(row) + strconv.Itoa(column) + strconv.Itoa(depth)
			}
		}
	}

	return cube
}

func arraySlicing() []string {
	myArray := [5]string{"iOS", "Android", "Windows", "Linux", "MacOS"}
	// the upper_indicator is where to start and lower_indicator is where to end
	// the results of slicing an array is a slice and slice is still a slice
	return myArray[4:] // slice is returned
}

func copyArrays() ([4]int, [4]int) {
	myArray := [4]int{1, 2, 3, 4}
	newArray := myArray
	newArray[0] = 0
	return newArray, myArray
}

func copySlices() ([]int, []int) {
	// inorder to copy a slice to another slice we use copy(destination, source), where destination is the slice that has a copy and source is the original
	// we use copy(), because slice is a pointer to the underlying array.
	mySlice := []int{1, 2, 3, 4}
	newSlice := make([]int, len(mySlice))

	copy(newSlice, mySlice)
	newSlice[0] = 0
	return mySlice, newSlice

}
