package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func main() {

	p := Pessoa{
		Nome:  "Ricardo",
		Idade: 19,
		Endereco: Endereco{
			Rua:    "B",
			Numero: 52,
			Cidade: "Sobradinho",
			Estado: "DF",
		},
	}
	imprimirpessoa(p)
}

func imprimirpessoa(p Pessoa) {
	endereco := p.Endereco

	fmt.Printf("Endereço de %s, %d anos:\n", p.Nome, p.Idade)
	fmt.Printf("Rua: %s, Número: %d\n", endereco.Rua, endereco.Numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", endereco.Cidade, endereco.Estado)

}
