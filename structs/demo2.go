package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("eklendi:", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Eylem", lastName: "Çiçek", age: 26}
	c.save()
	c.update()
}

func (c customer) update() {
	fmt.Println("güncenllendi:", c.firstName)
}
