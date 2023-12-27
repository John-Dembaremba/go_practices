package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// delay by 100 ms
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v) %v\n", i, s)
	}
}

func main() {
	go say("Hello", 3)
	go say("World", 2)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("============SYCHRONIZATION====================\n")
	execute()
	fmt.Printf("===========CHANNELS====================\n")
	channels()
}
