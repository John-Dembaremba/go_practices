package main

import (
	"fmt"
	"math"
)

type DigitCounter interface {
	CounterOddEven() (int int)
}

type DigitString string

func (ds DigitString) CounterOddEven() (int, int) {
	even, odd := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return odd, even
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	length float64
	name   string
}

// implement Shape Area to Circle
func (c Circle) Area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

// implement method not in Shape interface
func (c Circle) Circumeference() float64 {
	area := 2 * math.Pi * math.Pow(c.radius, 2)
	return area
}

// implement Shape Area to Square
func (ln Square) Area() float64 {
	area := math.Pow(ln.length, 2)
	return area
}

func main() {
	s := DigitString("123456")
	fmt.Println(s.CounterOddEven())
	fmt.Println()
	c1 := Circle{radius: 4, name: "circle"}
	sq := Square{length: 4, name: "square"}

	fmt.Printf("Areas of \nCircle: %v\n Square: %v\n Circumeference of Circle: %v\n", c1.Area(), sq.Area(), c1.Circumeference())
}
