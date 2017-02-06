package main

import (
	"fmt"
)

type Greeting struct {
	Name string
	Sex  string
	Age  int
}

func (g Greeting) SayHello1() {
	fmt.Printf("Changing name from %s to %s\n", g.Name, "NewName")
	g.Name = "NewName"
	fmt.Printf("New name: %s\n", g.Name)
	fmt.Printf("Hello %s (%s, %v)\n", g.Name, g.Sex, g.Age)
}

func (g *Greeting) SayHello2() {
	fmt.Printf("Changing name from %s to %s\n", g.Name, "NewName")
	g.Name = "NewName"
	fmt.Printf("New name: %s\n", g.Name)
	fmt.Printf("Hello %s (%s, %v)\n", g.Name, g.Sex, g.Age)
}

func main() {
	g1 := Greeting{
		"P",
		"M",
		22,
	}

	g1.SayHello1()
	// g.Name was changed inside SayHello but that does not reflect in g outside SayHello
	// This is because the receiver is not a pointer receiver. It's a value.
	fmt.Println(g1)

	// Now we will create a pointer to the Greeting struct
	g2 := &Greeting{
		"P",
		"M",
		22,
	}

	g2.SayHello2()
	// changing g.Name inside SayHello2 will change the original value of g.Name
	// because the receiver for SayHello2 is a pointer receiver and hence points
	// to the address of g2
	fmt.Println(g2)
}
