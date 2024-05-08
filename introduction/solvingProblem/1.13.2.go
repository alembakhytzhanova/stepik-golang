package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	res := ""
	for _, ch := range num {
		res = string(ch) + res
	}
	fmt.Println(res)
}
