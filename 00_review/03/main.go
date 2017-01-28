package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintf("This is sedan: %+v\n", s)
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprintf("This is truck: %+v\n", t)
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	fmt.Println("######## Cars ########")

	sedan1 := sedan{
		vehicle{3,
			"red"},
		false,
	}

	fmt.Println(sedan1)
	fmt.Println(sedan1.color)

	truck1 := truck{
		vehicle{3,
			"red"},
		false,
	}
	fmt.Println(truck1)
	fmt.Println(truck1.fourWheel)

	// fmt.Println(sedan1.transportationDevice())
	// fmt.Println(truck1.transportationDevice())

	report(sedan1)
	report(truck1)
}
