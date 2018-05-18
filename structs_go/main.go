package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // equal contactInfo
}

func main() {
	// kobe := person{firstName: "Kobe", lastName: "Bryant"}
	// kobe := person{"Kobe", "Bryant"}
	// fmt.Println(kobe.firstName)

	// var alex person = {"Alex", "Anderson"} is not work
	// var alex = person{"Alex", "Anderson"}
	// fmt.Println(alex)
	// fmt.Println("%+v", alex) // will not work as expected
	// fmt.Printf("%+v", alex) // %+v will print out all the properties in alex

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	// jane := person{
	// 	firstName: "Jane",
	// 	lastName:  "Forster",
	// 	contactInfo: contactInfo{
	// 		email:   "jane@findata.com.tw",
	// 		zipCode: 94000,
	// 	},
	// }

	// fmt.Printf("%+v", jane)

	// &variable is a variable operator,
	// meaning give me the memory address of the value this variable is pointing at
	// janePointer := &jane
	// jane.print()
	// janePointer.changeName("june")
	// jane.print()
	// jane.changeName("Jonny") // also work because go will do the get pointer address for you
	// jane.print()

	mySlice := []string{"hi", "how", "are", "you"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// pointerToPerson is an operator,
// means we want to manipulate the value the pointer is referencing
// *person is a type description, means we are working with a pointer to person
func (pointerToPerson *person) changeName(newFirstName string) {
	// *pointer is pointer...,
	// meaning give me a value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "bye"
}

// when we define a slice go will automatic create a slice and array for us
// slice contain the three element 1. pointer to head over 2. the capacity of slice 3. the length of slice
// and the array contain the value of each element in slice
// so actually is like we go in to slice and fetch some information then we get value in array
