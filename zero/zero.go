package main

import "fmt"

func main() {
	// zero value
	var str string
	var boolean bool
	var number int
	var decimal float32

	fmt.Println("default string:", str)
	fmt.Println("default bool:", boolean)
	fmt.Println("default non-decimal:", number)
	fmt.Println("default decimal:", decimal)
}
