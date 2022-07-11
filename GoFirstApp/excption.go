package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//nil = null ablamına gelir GO da
	//err kullanmak istemiyorsak _ kullanılabilir
	//file, _ := os.Open("file.txt")
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("hata:", err)
	}

	fmt.Println(file)

	//Kendimiz hata fırlatmak istersek bu şekilde yapıyoruz
	myErr := errors.New("it is a error")

	fmt.Println(myErr)
}
