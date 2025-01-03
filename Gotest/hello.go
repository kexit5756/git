package main

import (
	"GITTEST/GoTest/Tool"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, Go!")
	fmt.Println("My first Go program.")
	fmt.Println(add(1, 3))
	i := Tool.Decline(10, 3)
	fmt.Println(i)

}
