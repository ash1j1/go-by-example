package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// Won't work cause Go requires explicit type information 
	// for multi-variable declarations without assignment.
	// var num, numm = 10, 19

	var first_name = "John"
	var last_name = "Doe"
	fmt.Println(first_name + " " + last_name)
	fmt.Println(first_name + " Doe")

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f:= "apple"
	fmt.Println(f)
}
