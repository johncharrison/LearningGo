package main

import "fmt"

func main() {
	// Define a map

	// emails := make(map[string]string)

	// // Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@yahoo.com"
	// emails["Jim"] = "jim@gmail.com"

	// fmt.Println(emails)
	// fmt.Println(len(emails))
	// fmt.Println(emails["Bob"])

	// delete(emails, "Bob")
	// fmt.Println(emails)

	emails := map[string]string{"Bob": "bob@bob.com", "Sharon": "sharon@sharon.com"}
	fmt.Println(emails)
}
