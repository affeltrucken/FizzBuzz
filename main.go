package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter iterations for fizzbuzz: ")
	var iters int
	fmt.Scanln(&iters)
	fizzBuzz(iters)
}

func fizzBuzz(iters int) {
	for i := 1; i < iters; i++ {
		fmt.Printf("%d : ", i)
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Print("\n")
	}
}
