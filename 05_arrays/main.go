package main

import "log"

func main() {
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// Slices
	fruitArr := []string{"Apple", "Orange", "Grape", "Cucumber"}

	// [Orange Grape]
	log.Println(fruitArr[1:3])
}
