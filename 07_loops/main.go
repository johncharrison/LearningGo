package main

import (
	"fmt"
)

func main() {
	// var i int = 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i+1
	// 	// i++
	// 	i += 1
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i%5 == 0 {
			fmt.Printf("Buzz\n")
		} else if i%3 == 0 {
			fmt.Printf("Fizz\n")
		} else {
			fmt.Println(i)
		}
	}

}
