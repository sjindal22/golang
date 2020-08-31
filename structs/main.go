package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// shiv := person{"Shivika", "Jindal"}
	// shiv := person{firstName: "Shivika", lastName: "Jindal"}

	shiv := person{
		firstName: "Shivika",
		lastName:  "Jindal",
		contact: contactInfo{
			email:   "shivka.jindal@gmail.com",
			zipCode: "94000",
		},
	}

	fmt.Println(shiv)

	// defining a variable of type person, which is of type struct
	var rish person

	// GO by default assigns zero values to the variables that are not initialized
	// Output: {firstName: lastName:}
	fmt.Printf("%+v\n", rish)

	// Updating values in struct
	shiv.lastName = "Jain"
	fmt.Println(shiv)

	// This is another way to define a pointer
	//shivPointer := &shiv
	shiv.updateFirstname("Shiv")
	shiv.print()

}

func (pointerToValue *person) updateFirstname(newName string) {
	(*pointerToValue).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
