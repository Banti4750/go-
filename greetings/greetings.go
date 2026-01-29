package main

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	return fmt.Sprintf("Hi, %v. Welcome!!!!!!", name)
}

func main() {
	fmt.Println(Hello("Banti Kumar"))
}
