package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println()
	aFruits := fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	fmt.Println()
	bFruits := append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape", "banana", "papaya"]

	fmt.Println()
	dst := make([]string, 3)
	n := copy(dst, fruits)
	fmt.Println(fruits)
	fmt.Println(dst)
	fmt.Println(n)
}
