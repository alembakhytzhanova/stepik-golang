package main

import "fmt"

func main() {

	// "=="
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c) // false

}

func v() {
	// > <
	var a int = 8
	var b int = 3
	var c bool = a > b
	var g bool = a < b
	fmt.Println(c) // true
	fmt.Println(g) // false

}


func g() {
	// ">="   "<="

	var a int = 8
	var b int = 3
	var c bool = a>= b 
	var g bool = a <= b
	fmt.Println(c) // true
	fmt.Println(g) // false
}


func c() {
	// " != "

	var a int = 8
	var b int = 3
	var c bool = a != b // true
	var d bool = a != 8 // false
	fmt.Println(c)
	fmt.Println(d)

}
