package main

import "log"

func double(x int) int {
	return 2 * x
}

func main() {
	x := 5
	log.Printf("%d doubled is %d\n", x, double(x))
}
