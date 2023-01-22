package main

import "fmt"

type book struct {
	name   string
	writer string
	page   int
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

}
