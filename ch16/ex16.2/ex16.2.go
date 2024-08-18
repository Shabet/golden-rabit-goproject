package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	sliece2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(sliece2)
}
