package main

import "fmt"

func main() {
	for i, v := range []int{1, 23, 4, 5, 6} {
		fmt.Println(i, v)
	}

}
