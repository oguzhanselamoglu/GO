package main

import "fmt"

func main() {
	arr := [3]int{}

	arr[0] = 1
	arr[1] = 29
	arr[2] = 10

	fmt.Println(arr)

	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{4, 5, 2, 7, 9}

	fmt.Println(numbers)
	fmt.Println("number len", len(numbers))

	var myNumbers = [...]int{4, 5, 2, 7, 9, 43, 65}
	fmt.Println(myNumbers)
	fmt.Println("number len", len(myNumbers))
}
