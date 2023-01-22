package main

import (
	"errors"
	"fmt"
)

func checkUserName(v string) error {
	if v == "oguzhan" {
		return errors.New("Bu kullanıcı alınmış")
	}
	return nil
}
func main() {

	userNames := []string{"ali", "oguzhan"}
	for _, i := range userNames {
		if a := checkUserName(i); a != nil {
			fmt.Println(a)
		} else {
			fmt.Println(i, "Kullanıcı adı uygun")
		}
	}
}
