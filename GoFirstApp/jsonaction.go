package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

func main() {
	a := map[string]int{"ali": 40, "veli": 50}
	j, _ := json.Marshal(a)

	fmt.Println(string(j))
	b := []byte(string(j))

	var data map[string]int
	err := json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	user := User{
		Name:  "oguzhan",
		Email: "oguzhanselamoglu@gmail.com",
	}

	c, _ := json.Marshal(user)
	fmt.Println(string(c))

}
