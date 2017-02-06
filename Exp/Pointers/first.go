package main

import (
	"log"
)

type Person struct {
	name       string
	nextPerson *Person
}

func main() {
	first_person := Person{name: "First"}
	second_person := Person{name: "Second"}
	first_person.nextPerson = &second_person

	log.Println(*first_person.nextPerson)
	log.Println(first_person.nextPerson)
}
