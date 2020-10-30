package main

import "fmt"

func main()  {
	i := 1
	var p *int = nill

	p=&i
	*p++
	i++

	fmt.println(p, *p, &i)
	
}