package main

import "fmt"

func main() {
	// aritmatika
	value := (((2+6)%3)*4 - 2) / 3
	fmt.Println("value: ", value) // 2
	fmt.Println()

	// perbandingan
	fmt.Println("value = 2:", value == 2) // true
	fmt.Println("value > 0:", value > 0)  // true
	fmt.Println("value < 0:", value < 0)  // false
	fmt.Println()

	// logika
	benar := true
	salah := false
	fmt.Println("AND:", benar && salah) // false
	fmt.Println("OR:", benar || salah)  // true
	fmt.Println("Negasi:", !salah)      // true
}
