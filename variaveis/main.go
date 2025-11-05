package main

import (
	"fmt"
	variaveis "variaveis/tipagens"
)

const TaxaConversao = 5.25
const Moeda = "USD"

func main() {
	reais := 250.0
	dolar := reais / TaxaConversao
	fmt.Printf("%.2f BRL equivalem a %.2f USD", reais, dolar)
	fmt.Println()
	variaveis.Notas()
}
