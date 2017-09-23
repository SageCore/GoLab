//Run code
package main

import "fmt"

//This person struct type has name and age fields.
type car struct {
	name  string
	Model int
}

func main() {

	//This syntax creates a new struct.
	fmt.Println(car{"Odern Captima", 1950})
	//You can name the fields when initializing a struct.
	fmt.Println(car{name: "Honda civic", Model: 2009})
	//Omitted fields will be zero-valued.
	fmt.Println(car{name: "Lamborghini Aventador"})
	//An & prefix yields a pointer to the struct.
	fmt.Println(&car{name: "Honda city ", Model: 2016})
	//Access struct fields with a dot.
	s := car{name: "Mustang Gt", Model: 2010}
	fmt.Println(s.name)
	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.Model)
	//Structs are mutable.
	sp.Model = 2015
	fmt.Println(sp.Model)
}
