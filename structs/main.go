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

	var juca person
	juca.firstName = "Juca"
	juca.lastName = "Feliz"

	fmt.Println(thyago)
	//fmt.Println(john)
	fmt.Printf("%+v", juca)

}
