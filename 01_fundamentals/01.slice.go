package main

import "fmt"

var (
	Str []string
	l   string
)

func main() {
	fmt.Println(Str)
	Str := append(Str, "Hello all\n")
	fmt.Println(Str)
	Str = append(Str, "My name is Kaido\n")
	fmt.Println(Str)

	for _, v := range Str {
		l := l + v
		fmt.Print(l)
	}
}
