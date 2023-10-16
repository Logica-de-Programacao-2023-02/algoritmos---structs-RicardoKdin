package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func main() {
	animal := Animal{
		Nome:    "Galinha",
		Especie: "Frango",
		Idade:   10,
		Som:     "Pó",
	}

	AnimalInfo(animal)

	animal = AlteraSom(animal, "Có")

	AnimalInfo(animal)
}

func AlteraSom(animal Animal, novoSom string) Animal {
	animal.Som = novoSom
	return animal
}

func AnimalInfo(banana Animal) {
	fmt.Printf("O animal %s da espécie %s e com a idade de %d faz o seguinte som: %s\n",
		banana.Nome,
		banana.Especie,
		banana.Idade,
		banana.Som,
	)
}
