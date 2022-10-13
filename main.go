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
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
