package main

import "fmt"

type icontact interface {
	save()
	update()
	delete()
}
type contact struct {
	name   string
	number string
}

func (r contact) delete() {
	fmt.Println(r.name, "delete")
}

func (r contact) save() {
	fmt.Println(r.name, "save")
}
func (r contact) update() {
	fmt.Println(r.name, "update")
}

func action(r icontact) {
	r.save()
}
func main() {
	person := contact{
		name:   "Ali",
		number: "111111",
	}
	fmt.Println(person)
	action(person)

}
