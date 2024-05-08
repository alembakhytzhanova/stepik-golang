package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}

	}

	//fmt.Println(newArray)
}
