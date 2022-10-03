package main

import "fmt"

func main() {
	// tipe data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	message = `Nama saya "Ghifari".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)
}
