package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

	var green uint8 = 4
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

}
