package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	john := person{"John", "Travolta"}
	thyago := person{firstName: "Thyago", lastName: "Volpatto"}

	fmt.Println(thyago)
	fmt.Println(alex)

}
