package main

import "fmt"

func main() {
	// 3 rows, 4 columns, initialized with all zeroes:
	var a [3][4]int

	// Display entire array:
	fmt.Printf("%v\n", a)

	// Single element (Go arrays are 0-based):
	fmt.Printf("Row 0, column 0 is %v\n", a[0][0])

	a[0][0] = 10
	a[1][0] = 20

	// Single elements, with values assigned:
	fmt.Printf("Row 0, column 0 is %v\n", a[0][0])
	fmt.Printf("Row 1, column 0 is %v\n", a[1][0])

	// Display entire array:
	fmt.Printf("%v\n", a)
}
