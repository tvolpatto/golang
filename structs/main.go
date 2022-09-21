package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//john := person{"John", "Travolta"}
	thyago := person{
		firstName: "Thyago",
		lastName:  "Volpatto",
		contact: contactInfo{
			email:   "myemail@gmail.com",
			zipcode: 90887,
		},
	}

	var juca person
	juca.firstName = "Juca"
	juca.lastName = "Feliz"

	fmt.Println(thyago)
	//fmt.Println(john)
	fmt.Printf("%+v", juca)

}
