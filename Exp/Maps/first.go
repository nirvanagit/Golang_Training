package main

import (
	"log"
)

func main() {
	m := make(map[string]int)
	b := 2
	s := []string{"a", "b"}
	m["a"] = 1
	m["b"] = 2
	log.Printf("value of the map: %s", m)
	log.Printf("value of b: %d", b)
	log.Printf("value of s: %v", s)
	changeMap(m, b, s...)
	log.Printf("value of the map: %s", m)
	log.Printf("value of b: %d", b)
	log.Printf("value of s: %v", s)
}

func changeMap(m map[string]int, d int, s []string) {
	m["c"] = 3
	d = 5
	s = []string{"c", "d"}
}
