package main

import "fmt"

func main() {
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}

func c() {
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
}

// Вывод:
// 42
// 1
// default

func a() {
	var c uint32
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("от 1 до 9")
	case 100 <= c && c <= 250:
		fmt.Println("от 100 до 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("от 1000 до 6000")
	}
}
