package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type Maquina struct {
	OpcaoDeEscolha OpcaoDeEscolha
}

func NovaMaquina() Maquina {

	x := rand.Intn(3)
	fmt.Println(x)

	return Maquina{
		OpcaoDeEscolha: OpcoesDeEscolha()[x],
	}
}
