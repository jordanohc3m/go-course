package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //(float64 já inferido pelo compilador)

	//forma reduzida com ":="
	area := PI * math.Pow(raio, 2)

	fmt.Println("Área da circunferência:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "LAU!"

	fmt.Println(g, h, i)

}
