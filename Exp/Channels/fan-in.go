package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Multiple channels should write to a channel.
	out := fanIn(tring("Anubhav"), tring("Puja"))

	// Then print the values inside the channel which contains values
	// from multiple channels
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}

	fmt.Println("It's time for a break")
}

func tring(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s says hello there!", name)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(i1, i2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-i1
		}
	}()

	go func() {
		for {
			c <- <-i2
		}
	}()
	return c
}
