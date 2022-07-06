package main

import "fmt"

var (
	d1 = "degisken"
	d2 = 5
)

func main() {
	var a string = "değer"
	var b = "string"
	c := "string 3"
	var d bool = true
	var e int = 10
	var f int = 20
	var aa, bb, cc int = 0, 1, 2

	var aaa, bbb, ccc = 0, true, 4.3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("e + f = ", e+f)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)

}
