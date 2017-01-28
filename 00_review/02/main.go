package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	fmt.Println("######## Person #######")

	p1 := person{
		fName:   "Radek",
		lName:   "Sohlich",
		favFood: []string{"pizza", "burger"},
	}

	fmt.Println(p1)
	fmt.Println(p1.fName)
	for idx, val := range p1.favFood {
		fmt.Printf("index: %d value: %d\n", idx, val)
	}
}
