package main

import "fmt"

func main() {
	fruits := [...]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Println("nama buah:", fruit)
	}
	fmt.Println()
	for i := range fruits {
		fmt.Println("index:", i)
	}
}
