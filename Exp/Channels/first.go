package main

import (
	"log"
)

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// This is blocking code. As the statement range c is
	// waiting to receive any value inside c
	for n := range c {
		log.Println(n)
	}
}
