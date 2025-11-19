package main

import "fmt"

type Lista struct {
	Valor int
	Prox  *Lista
}

func AdicionaNo(no *Lista, valor int) *Lista {
	if no == nil {
		return &Lista{Valor: valor}
	}
	if no.Prox == nil {
		no.Prox = &Lista{Valor: valor}
	} else {
		no.Prox = AdicionaNo(no.Prox, valor)
	}
	return no
}
func RemoveNo(no *Lista, valor int) *Lista {
	if no == nil {
		return nil
	}
	if no.Valor == valor {
		return no.Prox
	}
	no.Prox = RemoveNo(no.Prox, valor)
	return no
}

func ImprimeLista(no *Lista) {
	if no != nil {
		fmt.Println(no.Valor)
		if no.Prox != nil {
			ImprimeLista(no.Prox)

		}
	}
}
func main() {
	var l1 *Lista = nil

	ImprimeLista(l1)

	l1 = AdicionaNo(l1, 10)
	l1 = AdicionaNo(l1, 20)
	l1 = AdicionaNo(l1, 30)
	l1 = AdicionaNo(l1, 40)
	l1 = AdicionaNo(l1, 50)
	l1 = AdicionaNo(l1, 60)
	l1 = AdicionaNo(l1, 70)
	ImprimeLista(l1)

	RemoveNo(l1, 50)
	ImprimeLista(l1)
}
