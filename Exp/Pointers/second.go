package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	u := &User{Name: "Leto"}
	fmt.Println(u.Name)
	Modify(u)
	fmt.Println(u.Name)
}

func Modify(u *User) {
	u.Name = "Paul"
	// The below statement does not work,
	// because here we are reassigning the value of u.
	// u = &User{Name: "Paul"}
}
