package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
	*/
	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	goRoot := os.Getenv("GOROOT")
	fmt.Println("Username", uName)
	fmt.Println("uDomain", uDomain)
	fmt.Println("processorIdentifier", processorIdentifier)
	fmt.Println("goRoot", goRoot)
}
