package main

import "fmt"

func main() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a меньше, чем b")
	}
}

func v() {
	// "else" "else if"
	if a < b {
		fmt.Println("a smaller than b")
	} else if a > b {
		fmt.Println("a больше b")
	}
}

func g() {
	/* If we need an option when none
	of the conditions are met, then we use the else statement:
	*/
	if a < b {
		fmt.Println("a меньше b")
	}else if a > b {
		fmt.Println("a больше b")
	}else  {
		fmt.Println("a равно b")
	}
}
