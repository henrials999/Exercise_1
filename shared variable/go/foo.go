package main

import (
	"fmt"
	"runtime"
)

var i = 0

func incrementing(ch chan int, done chan bool) {
	for j := 0; j < 1000000; j++ {
		val := <-ch
		val++
		ch <- val
	}
	done <- true
}

func decrementing(ch chan int, done chan bool) {
	for j := 0; j < 999999; j++ {
		val := <-ch
		val--
		ch <- val
	}
	done <- true
}

func main() {
	runtime.GOMAXPROCS(2)

	ch := make(chan int, 1)
	done := make(chan bool)

	ch <- i

	go incrementing(ch, done)
	go decrementing(ch, done)

	<-done
	<-done

	i := <-ch

	fmt.Println("The magic number is:", i)
}
