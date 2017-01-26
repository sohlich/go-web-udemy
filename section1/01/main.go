package main

import "fmt"

func main() {

	fmt.Println("###### Slice ########")
	slice := []int{5, 6, 7}

	fmt.Println(slice)

	for idx, val := range slice {
		fmt.Printf("index: %d value: %d\n", idx, val)
	}

	fmt.Println("######## Map ##########")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(m)

	for key, val := range m {
		fmt.Printf("key: %s value: %d\n", key, val)
	}
}
