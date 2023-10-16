package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {

	viagens := []Viagem{
		{Origem: "Cidade A", Destino: "Cidade B", Data: "2023-10-20", Preco: 500.0},
		{Origem: "Cidade C", Destino: "Cidade D", Data: "2023-10-15", Preco: 700.0},
		{Origem: "Cidade E", Destino: "Cidade F", Data: "2023-10-25", Preco: 600.0},
	}

	viagemCara := viagemMaisCara(viagens)

	fmt.Printf("A viagem mais cara é de %s para %s, com data em %s e preço de R$ %.2f\n", viagemCara.Origem, viagemCara.Destino, viagemCara.Data, viagemCara.Preco)
}
