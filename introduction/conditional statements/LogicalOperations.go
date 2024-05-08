package main

func a() {
	// "!"
	var a bool = true
	var b bool = !a // false
	var c bool = !b // true

	// "&&"
	var b bool = 4 > 5 && 6 > 8   // false
	var c bool = 3 <= 5 && 10 > 8 // true

	// "||"
	var b bool = 4 > 5 || 6 > 8   //false
	var c bool = 3 == 5 || 10 > 8 // true

}
