package main
import "fmt"

func main() {
    
	count := 0
     var size int
    fmt.Scan(&size)
    arr := make([]int, size)
    for i := 0; i < len(arr); i++ {
        fmt.Scan(&arr[i])
    }

	for _, ch := range arr {
		if ch >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
 