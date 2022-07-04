package main

import "log"

func main() {

	x := 15
	y := 15

	// If Else
	if x <= y {
		log.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		log.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "yellow"

	if color == "red" {
		log.Println("color is red")
	} else if color == "blue" {
		log.Println("color is blue")
	} else {
		log.Println("color is not blue or red")
	}

	switch color {
	case "red":
		log.Println("color is red")
	case "blue":
		log.Println("color is blue")
	default:
		log.Println("color is not blue or red")
	}
}
