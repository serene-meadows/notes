package main

import "fmt"

func main() {
	a := [10 * 1024 * 1024]byte{}
	fmt.Println(&a)
	//
	// b := [10*1024*1024 + 1]byte{}
	// fmt.Println(&b)
}
