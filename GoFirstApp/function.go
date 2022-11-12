package main

import "fmt"

func topla(a int, b int) int {
	return a + b
}

func islenler(a, b int) (int, int, int, int) {
	toplam := a + b
	carp := a * b
	bol := a / b
	cikar := a - b
	return toplam, carp, bol, cikar
}
func main() {
	t, c, b, d := islenler(8, 2)
	_, e, _, _ := islenler(8, 2)
	fmt.Println(topla(2, 5))
	fmt.Println("Toplam ", t)
	fmt.Println("Carpma ", c)
	fmt.Println("Bolme ", b)
	fmt.Println("Çıkarma ", d)
	fmt.Println("Sonuc ", e)

}
