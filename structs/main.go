package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	//john := person{"John", "Travolta"}
	thyago := person{
		firstName: "Thyago",
		lastName:  "Volpatto",
		contactInfo: contactInfo{
			email:   "myemail@gmail.com",
			zipcode: 90887,
		},
	}

	//& - get the memory address of the variable
	thPointer := &thyago
	thPointer.updateName("John")
	thyago.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// * - is a type description - it means we are working with a pointer
func (pointerToPerson *person) updateName(newFirstName string) {
	//* - give the value that memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}
