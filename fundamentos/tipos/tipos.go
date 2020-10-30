package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//só positivos uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Literal inteiro é", reflect.TypeOf(b))

	//sem sinal (positivos e negativos) uint8 uint16 uint32 uint64

	i1 := math.MaxInt64
	fmt.Println("Literal inteiro é", reflect.TypeOf(i1))

}
