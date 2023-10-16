package main

import (
	"fmt"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(percentual float64) {
	f.Salario += f.Salario * (percentual / 100)
}

func (f *Funcionario) DiminuirSalario(percentual float64) {
	f.Salario -= f.Salario * (percentual / 100)
}

func (f Funcionario) TempoServico() int {
	idadeAtual := f.Idade
	tempoServico := idadeAtual - 18
	return tempoServico
}

func main() {

	funcionario := Funcionario{
		Nome:    "Ricardo",
		Salario: 8000,
		Idade:   38,
	}

	fmt.Printf("Funcionário: %s\n", funcionario.Nome)
	fmt.Printf("Salário: R$ %.2f\n", funcionario.Salario)
	fmt.Printf("Idade: %d anos\n", funcionario.Idade)

	funcionario.AumentarSalario(10)
	fmt.Printf("Salário após aumento: R$ %.2f\n", funcionario.Salario)

	funcionario.DiminuirSalario(5)
	fmt.Printf("Salário após diminuição: R$ %.2f\n", funcionario.Salario)

	tempoDeServico := funcionario.TempoServico()
	fmt.Printf("Tempo de serviço: %d anos\n", tempoDeServico)
}
