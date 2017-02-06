package main

import (
	"log"
)

func main() {
	num := 5
	c := factorial(num)

	for n := range c {
		log.Print(n)
	}
}

func factorial(num int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
