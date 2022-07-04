package main

import "fmt"

func changeName(s *string) {
	*s = "Fred"
}

func main() {

	a := 5
	b := &a

	fmt.Println(*&a, b)
	fmt.Printf("%T\n", b)

	// Use * to read value from address (&)
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer

	*b = 10
	fmt.Println(a)

	name := "Jim"
	fmt.Println(name)
	changeName(&name)
	fmt.Println(name)
}
