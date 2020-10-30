package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	// int para string
	fmt.Println("Tete", strconv.Itoa(123))

	//string para int (_ para ignorar o segundo retorno da func, no casso um erro)
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println(b)

	}

}
