package main

import "fmt"

type Flamengo struct {
	Brasileirao  int
	CopaBR       int
	Libertadores int
}

func (mengo Flamengo) Apresentar() {
	fmt.Printf("Tem %d Campeonatos Brasileiros, %d Copas do Brasil e %d Libertadores\n", mengo.Brasileirao, mengo.CopaBR, mengo.Libertadores)
}

func main() {
	f1 := Flamengo{Brasileirao: 8, CopaBR: 5, Libertadores: 3}
	f1.Apresentar()
}
