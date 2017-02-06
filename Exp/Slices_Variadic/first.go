package main

import (
	"fmt"
)

func main() {
	// a := []string{"a", "b"}
	printIt()
}

func printIt(a ...string) {
	for _, v := range a {
		fmt.Printf("%s\n", v)
	}
}
