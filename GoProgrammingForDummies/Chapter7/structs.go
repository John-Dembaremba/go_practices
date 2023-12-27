package main

import (
	"fmt"
)

type fieldPoints struct {
	x float64
	y float64
	z float64
}

type Pointer struct {
	X    float64
	Y    float64
	Z    float64
	Name []string
}

func main() {

	fpt1 := fieldPoints{x: 1.3, y: 2.5, z: 3.0}
	fpt2 := fieldPoints{x: 12.33, y: 892.112}
	fpt3 := structWithPointer(1000.9, 188.21, 950.76)

	fmt.Println(fpt1)
	fmt.Println()
	fmt.Println(fpt2)
	fmt.Println()
	fmt.Println(fpt2.x)
	fmt.Println()
	fmt.Println(fpt1.x)
	fmt.Println()
	fmt.Printf("##### FPT3 ########\n")

	fmt.Println(fpt3)
	fmt.Printf("\n================\n")

	// Copying structs
	fpt4 := fpt3
	fpt5 := *fpt4
	fpt4.x = 23.333
	fmt.Printf("##### FPT4 ########\n")
	fmt.Print(fpt4)

	// Confirm that its a pointer
	fmt.Printf("\n##### FPT3 ########\n")

	fmt.Println(fpt3)

	// create independent copy of structs
	fmt.Printf("\n##### FPT5 ########\n")
	fmt.Print(fpt5)

	fmt.Printf("\n##### COMPARE ftp1 and ftp2 ########\n")
	fmt.Println(compareStructs(fpt1, fpt2))

	fmt.Printf("\n##### COMPARE pt1, pt2 and pt3 which has data structures ########\n")
	// cmp package is used for this type of comparisons.
	// pt1 := Pointer{X: 5.6, Y: 3.8, Z: 6.9,
	// 	Name: []string{"pt1"}}
	// pt2 := Pointer{X: 5.6, Y: 3.8, Z: 6.9,
	// 	Name: []string{"pt"}}
	// pt3 := Pointer{X: 5.6, Y: 3.8, Z: 6.9,
	// 	Name: []string{"pt"}}
	// fmt.Printf("##### pt1 == pt2 ########\n")
	// fmt.Println(cmp.Compare[Pointer](pt1, pt2))
}

func structWithPointer(_x, _y, _z float64) *fieldPoints {
	p := fieldPoints{x: _x, y: _y, z: _z}
	return &p
}

func compareStructs(p1, p2 fieldPoints) bool {
	return p1 == p2
}
