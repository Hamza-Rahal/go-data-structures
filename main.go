package main

import (
	"fmt"

	"./stack"
)

func assert(predicate bool, message string) {
	if !predicate {
		panic(message)
	}
}

func main() {
	s := stack.New()
	s.Add(5)
	s.Add(6)
	s.Add(7)
	s.Add(8)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
