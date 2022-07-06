package main

import "fmt"

func main() {
	a := 10
	b := 20

	var total int = a + b
	fmt.Println("total :", total)

	total *= 2
	fmt.Println("total :", total)

	total -= 10

	fmt.Println("total :", total)

	total = b / a
	fmt.Println("total :", total)

	total++

	fmt.Println("total :", total)
}
