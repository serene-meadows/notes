package main

import "fmt"

func main() {
	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// using closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
}
