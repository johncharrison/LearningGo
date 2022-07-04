package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	//int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint 64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var

	var age byte = 28
	const isCool = true

	// shorthand
	name, email := "Jim", "jim@gmail.com"

	var pi float32 = 3.141592

	fmt.Println(name, email, age, isCool, pi)
	fmt.Printf("%T\n", pi)
}
