package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTvGrande := trab1 && trab2
	comprarTvPequena := trab1 != trab2 //ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTvGrande, comprarTvPequena, comprarSorvete

}

func main() {
	tvGrande, tvPequena, sorvere := compras(true, true)
	fmt.Printf("Tv Grande: %t, Tv Pequena: %t, Sorva: %t", tvGrande, tvPequena, !sorvere)

}
