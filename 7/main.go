package main

import "fmt"

var msg string

func init() {
	msg = "Hello, World!"
}

func main() {
	fmt.Println(msg)
	updateMsg(&msg)
	fmt.Println(msg)
}

// update msg by ref
func updateMsg(msg *string) {
	*msg = "Hello, World! Updated"
}
