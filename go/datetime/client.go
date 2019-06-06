package main

import (
	"fmt"
)

func main() {
	myDateTime := DateTime{
		InputDateTime: "2018-09-24T13:15:20-04:00"}
	myDateTime.ParseInput()

	if myDateTime.CurrentError != nil {
		fmt.Println(myDateTime.CurrentError)
	}

	fmt.Println("Month is " + myDateTime.Month)
	fmt.Println("Day is " + myDateTime.Day)
	fmt.Println("Year is " + myDateTime.Year)

	fmt.Println("Hour is " + myDateTime.Hour)
	fmt.Println("Minute is " + myDateTime.Minute)
	fmt.Println("Second is " + myDateTime.Second)

	fmt.Println("Formatted time is " + myDateTime.FormattedTime)

	fmt.Println("TimeZone is " + myDateTime.TimeZone)
}
