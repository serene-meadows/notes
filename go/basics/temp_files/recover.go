package main

import "fmt"

func main() {
	process()
	fmt.Println("return from main")
}

func process() {
	defer func() {
		fmt.Println("just checking")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("Start process")

	panic("panicking")

	// fmt.Println("End process")
	// commented because statements after panic() is called
	// are not reachable
}
