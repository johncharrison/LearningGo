package main

import "fmt"

func main() {
	ids := []int{32, 42, 55, 77, 86, 454, 7886}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not Using Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Sum ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "Tim": "tim@gmail.com", "Jen": "jen@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
