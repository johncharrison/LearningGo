package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(p2 *Person) {
	if p.gender == "M" {
		p2.lastName = p.lastName
	} else {
		p.lastName = p2.lastName
	}
}

func main() {
	// init Person using struct

	person1 := Person{
		firstName: "Jane",
		lastName:  "Smith",
		city:      "NYC",
		gender:    "F",
		age:       28,
	}

	person2 := Person{
		firstName: "Bob",
		lastName:  "Dole",
		city:      "Omaha",
		gender:    "M",
		age:       32,
	}

	fmt.Println(person1)
	person1.getMarried(&person2)
	greeting1 := person1.greet()
	greeting2 := person2.greet()

	fmt.Println(greeting1)
	fmt.Println(greeting2)

}
