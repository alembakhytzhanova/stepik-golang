package mai

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	digit1 := number / 100
	digit2 := (number % 100) / 10
	digit3 := number % 10

	sum := digit1 + digit2 + digit3

	fmt.Println(sum)
}
