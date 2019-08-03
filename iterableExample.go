package main

import "fmt"

type Iterable <-chan int

func main() {
	c := make(chan int)
	go func() {
		c <- 8
		c <- 9
	}()

	iter := Iterable(c)

	fmt.Println(<-iter)
	fmt.Println(<-iter)
	fmt.Println(<-iter)

}
