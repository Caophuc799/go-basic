package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
	string
}

func (p person) print() {
	fmt.Println("=============")
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string, newLastName string) {
	(*p).firstName = newFirstName
	p.lastName = newLastName
}

func main() {
	alex := person{"Cao", "Phuc", contactInfo{}}
	fmt.Println(alex)

	nick := person{firstName: "Cao", lastName: "Phuc"}
	fmt.Println(nick)

	var bart person
	fmt.Println(bart)
	bart.firstName = "Kim"
	bart.lastName = "Tiem"
	fmt.Printf("%+v", bart)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9000,
			string:  "helo",
		},
	}

	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmmmm", "Me")
	jim.print()

	jim.updateName("Jimmmmm", "Te")
	jim.print()
}
