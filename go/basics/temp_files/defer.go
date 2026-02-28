package main

import "fmt"

func main() {
	process(0)
}

func process(num int) {
	defer fmt.Println("deferred statement 1 execution", num)
	num++
	defer fmt.Println("deferred statement 2 execution", num)
	num++

	fmt.Println("normal statement execution", num)
}
