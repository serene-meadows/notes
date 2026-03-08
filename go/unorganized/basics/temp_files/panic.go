package main

import "fmt"

func main() {
	process(10)
	process(-1)
}

func process(input int) {
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")

	if input < 0 {
		fmt.Println("before panic")
		panic(1)
		// fmt.Println("after panic")
		// defer fmt.Println("deferred 2")
		// both statements are unrecheable
	}

	fmt.Println("processing input:", input)
}
