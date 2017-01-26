package main

import "fmt"

type gator int

func (g gator) greetings() {
	fmt.Println("Hi I'm gator.")
}

type flaminco bool

func (f flaminco) greetings() {
	fmt.Println("Hi I'm flaminco.")
}

func main() {
	var g1 gator
	g1 = 1
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Printf("%T\n", x)

	x = int(g1) // Type assertion only for interfaces

	g1.greetings()
}
