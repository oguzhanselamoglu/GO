package main

import "fmt"

func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("büyüktür")
	} else if a == b {
		fmt.Println("eşittir")
	} else {
		fmt.Println("küçüktr")
	}

	if d := 10; d > 10 {
		fmt.Println("içeride")
	}

	c := 4
	switch {
	case c == 4:
		println("dört")
	case c > 5:
		println("beş")
	default:
		println("default")
	}
}
