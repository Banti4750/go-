package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b

}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Sum of 3 and 4 is:", sum(3, 4))
	fmt.Println("Difference of 7 and 5 is:", diff(7, 5))
}
