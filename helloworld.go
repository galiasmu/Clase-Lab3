package main

import (
	"fmt"
)

type operaciones struct {
	a int;
	b int;
}



func Suma(a int, b int ) int {
	return a+b
}

func Resta(a int, b int ) int {
	return a-b
}

func Producto(a int, b int ) int {
	return a*b
}

func Division(a int, b int ) int {
	return a/b
}

func main() {
	fmt.Println(Suma(2,3))
	fmt.Println(Resta(2,3))
	fmt.Println(Producto(2,3))
	fmt.Println(Division(2,3))
}