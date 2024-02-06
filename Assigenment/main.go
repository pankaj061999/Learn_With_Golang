package main

import "fmt"

func main() {
	var userSlice = []int64{0, 1, 3, 5, 2, 6, 7, 9}

	for i := 0; i < len(userSlice); i++ {
		fmt.Println("Print i", i)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("odd")
		}
	}
}
