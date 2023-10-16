package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverUltimaNota() {
	if len(a.Notas) > 0 {
		a.Notas = a.Notas[:len(a.Notas)-1]
	}
}

func (a Aluno) CalcularMedia() float64 {

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}

	return soma / float64(len(a.Notas))
}

func (a Aluno) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d anos\n", a.Idade)
	fmt.Printf("Média das notas: %.2f\n", a.CalcularMedia())
}

func main() {

	aluno := Aluno{
		Nome:  "Ricardo",
		Idade: 19,
		Notas: []float64{8.0, 9.0, 7.0},
	}

	aluno.ImprimirInformacoes()

	aluno.RemoverUltimaNota()

	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(10.0)

	aluno.ImprimirInformacoes()
}
