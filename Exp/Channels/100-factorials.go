package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		for n := range factorial(5) {
			fmt.Printf("Result of %vth computation : %v\n", i, n)
		}
	}
}

func factorial(num int) chan int {
	c := make(chan int)

	go func() {
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
