package main

import "fmt"

func main() {
	ch := make(chan int, 1_000_000)
	for num := range 1_000_000 {
		go func(num int) {
			fmt.Println("running", num)
			ch <- num
		}(num)
	}

	for range 1_000_000 {
		<-ch
	}
}
