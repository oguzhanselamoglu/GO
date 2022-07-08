package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s = "1"
	var x = 10
	var f float32 = 2.0

	fmt.Println(s, x, f)
	number, _ := strconv.Atoi(s)
	fmt.Println(number)
	//Atoi : string to int
	//Itoi :int to string

	fmt.Println("sonuc:" + strconv.Itoa(number))

	//veri tipini parantez ile ifade edersek casting işlemi yapar
	//küçükten büyüğe atama işlemi gerçekleşir
	var i int = 55
	var c float64 = float64(i)
	fmt.Println(i, c)

}
