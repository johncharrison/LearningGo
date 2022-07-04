package main

import (
	"fmt"
	"math"
	"strutil"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse(strutil.Reverse("Hola")))
}
