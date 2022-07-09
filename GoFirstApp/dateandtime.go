package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, time.November, 6, 4, 0, 0, 0, time.UTC)

	fmt.Printf("Go lunch at %s\n", t)
	fmt.Println("----------------------")

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)
	fmt.Println("----------------------")

	fmt.Println("The month is %s\n", t.Month())
	fmt.Println("the day is %s\n", t.Day())
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
}
