package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}
type Endereco struct {
	Rua    string
	Numero int
	Bairro string
}

func main() {
	cliente1 := Cliente{
		Nome:  "Ueslei",
		Idade: 25,
		Endereco: Endereco{
			Rua:    "rua tal",
			Numero: 40000,
			Bairro: "bairro top",
		},
		Email: "ueslei392@gmail.com",
	}

	fmt.Println(cliente1)
	cliente1.Endereco.Numero = 30000
	fmt.Printf("Numero do cliente: %d", cliente1.Endereco.Numero)
}
