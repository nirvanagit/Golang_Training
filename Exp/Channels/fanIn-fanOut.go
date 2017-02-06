package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen(2, 3)

	// Fan Out
	// Give output from a channel to 2 or more functions
	c1 := sq(in)
	c2 := sq(in)

	// Fan In
	// Use the output from two different functions and
	// put it in channels.
	// Then a single channel is going to listen to the channels
	// above and provide the output
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, n := range cs {
		// Question: Why do we need a closure here?
		// Answer: This loop is going to create
		// go routines which are ran concurrently.
		// Lets say it loops twice. During the first
		// loop, the value inside n is 1. But by the time it reaches
		// to the line where it has to give the value to the channel, the second loop would
		// have run and would have changed the value of n to something else.
		// To make sure that out <- n gets the correct context we have to use
		// closures
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(n)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
