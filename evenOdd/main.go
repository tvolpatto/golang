package main

import "fmt"

func main() {
	numbers := startNumbers()

	fmt.Println("Even numbers in the selected interval: " + string(numbers.evens(10).toString()))
	fmt.Println("Odd numbers in the selected interval: " + string(numbers.odds(10).toString()))

}
