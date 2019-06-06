package main

import "fmt"

func sayHello() int {
	fmt.Printf("Hello, world!\n")

	return 0
}

func sayGoodbye() int {
	fmt.Printf("Goodbye, cruel world!\n")

	return 0
}

func main() {
	sayHello()

	sayGoodbye()
}
