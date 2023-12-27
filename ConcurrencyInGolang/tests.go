package main

import (
	"fmt"
	"time"
)

func main() {
	var data int

	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	for {
		if data == 0 {
			fmt.Println("Data zero ==>>", data)
			break
		} else {
			fmt.Println("Data zero ==>>", data)
		}
	}

}
