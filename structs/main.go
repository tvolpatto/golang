package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	john := person{"John", "Travolta"}
	thyago := person{firstName: "Thyago", lastName: "Volpatto"}

	var juca person
	juca.firstName = "Juca"
	juca.lastName = "Feliz"

	fmt.Println(thyago)
	fmt.Println(john)
	fmt.Printf("%+v", juca)

}
