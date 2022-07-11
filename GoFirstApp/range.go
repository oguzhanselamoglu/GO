package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// i index v ise değer
	for i, v := range pow {
		fmt.Printf("2+%d =%d\n", i, v)
	}

	a := [...]string{"a", "b", "c", "d", "e"}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	capitals := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Roma"}

	for key := range capitals {
		fmt.Println("Map item: capital of", key, "is", capitals[key])
	}
}
