package main

import "fmt"

func main() {
	var workArray [10]uint
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	for j := 0; j < 3; j++ {
		var index1, index2 int
		fmt.Scan(&index1, &index2)

		if index1 >= 0 && index1 < 10 && index2 >= 0 && index2 < 10 {
			workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
		}
	}

	for _, value := range workArray {
		fmt.Print(value, " ")
	}
	fmt.Printf("type ok")
	//return
}
