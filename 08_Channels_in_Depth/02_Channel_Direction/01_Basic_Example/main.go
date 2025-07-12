/*
Default channel is bidirectional, but you can create send-only or receive-only channels.
Send-only channel: `chan<- int`
Receive-only channel: `<-chan int`

*/

package main

import "fmt"

func sendData(ch chan<- int) { // send-only channel
	ch <- 100
	fmt.Println("Sent data")
}

func receiveData(ch <-chan int) { // receive-only channel
	val := <-ch
	fmt.Println("Received data:", val)
}

func main() {
	ch := make(chan int)

	go sendData(ch)
	go receiveData(ch)

	// Give goroutines time to run
	fmt.Scanln()
}
