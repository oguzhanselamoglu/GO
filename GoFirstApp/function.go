package main

import (
	"fmt"
)

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

func person(name string) func() {
	_name := name
	return func() {
		fmt.Println(_name)
	}
}

func person2(name string) func() string {
	_name := name
	persons := func() string {
		return "değer " + _name
	}
	return persons
}

func count() func() {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func count2() (first func(), second func() int) {
	n := 0
	first = func() {
		n++
		fmt.Println(n)
	}
	second = func() int {
		return n
	}
	return
}

func faktor(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktor(n-1)
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
	p := person("oguzhan")
	p()
	newPerson := person2("selamoglu")
	f := newPerson()

	fmt.Println(f)

	//cnt := count()
	first, second := count2()
	for i := 0; i < 20; i++ {
		first()
	}
	fmt.Println(second())
}
