package main

import "fmt"

func main() {
	greetingMessage := "Hello, World!"
	fmt.Println(greetingMessage)

	var name string
	var age int
	var isCool bool

	name = "John Doe"
	age = 25
	isCool = true
	fmt.Printf("%s, age: %d, isCool: %t\n", name, age, isCool)

	var first, second, third = "first", "third", "second"
	second, third = third, second
	first, _, third = "first word", "empty string", "third word"
	fmt.Println(first, second, third)
}
