package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	fmt.Printf("Sending data to a channel...\n")
	// time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Printf("Message retrieved from channel...\n")

}

func getData(ch chan string) {
	fmt.Printf("Receiving data from a channel...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Message is %v\n", <-ch)
}

func channels() {
	ch := make(chan string) // unbuffered channels
	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
}
