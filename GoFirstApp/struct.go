package main

import "fmt"

type book struct {
	name   string
	writer string
	page   int
}

type circle struct {
	radius int
}

func (a *circle) area() int {
	return a.radius * a.radius * 3
}
func main() {
	fmt.Println(book{
		name:   "Kuçük Ağa",
		writer: "Ömer Seyfettin",
		page:   100,
	})

	a := book{
		name:   "Savaş ve Barış",
		writer: "Tolstoy",
		page:   200,
	}
	fmt.Println(a.name)

	b := circle{radius: 4}
	fmt.Println(b.area())

}
