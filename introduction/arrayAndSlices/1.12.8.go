package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]
	for i := 0; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	fmt.Println(max)

}