package main

import "fmt"

func main() {
	value := "Catamandoo"

	// Take length of string with len.
	length := len(value)
	fmt.Println(length)
	k := fmt.Sprintf("Bearer %s", value)
	fmt.Println(k)
}
