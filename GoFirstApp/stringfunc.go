package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "10,11,12,13,14,15"
	b := strings.Split(a, ",")
	fmt.Println(b[0])

	c := "oguzhanselamoglu"
	d := strings.Contains(c, "han")
	fmt.Println(d)

	e := strings.Count(c, "a")
	fmt.Println(e)

	f := strings.Index(c, "h")
	fmt.Println(f)

	r := strings.Replace(c, "selamoglu", "burak", 1)
	fmt.Println(r)

}
