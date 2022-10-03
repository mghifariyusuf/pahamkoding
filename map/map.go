package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	delete(chicken, "maret")

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	value, isExist := chicken["juni"]
	fmt.Println(value, isExist)
	value, isExist = chicken["april"]
	fmt.Println(value, isExist)
}
