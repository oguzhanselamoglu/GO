package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id   int
	Name string
}

func main() {
	url := "http://localhost:8080/products"
	resp, err := http.Get(url)
	checkError(err)

	jsonData, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	var data []Product
	err = json.Unmarshal([]byte(jsonData), &data)
	checkError(err)

	fmt.Println(data)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
