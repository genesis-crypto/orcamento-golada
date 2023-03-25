package main

import (
	"fmt"
	"orcamento-golada/calculadora"
	"orcamento-golada/orcamento"
)

func main() {
	calculadoraConcreta := &calculadora.Calculadora{}

	calculadoraICMS := &calculadora.CalculadoraICMS{
		Calculadora: calculadoraConcreta,
	}

	calculadoraPIS := &calculadora.CalculadoraPIS{
		Calculadora: calculadoraICMS,
	}

	produtos := []orcamento.Produto{
		{Nome: "Notebook", Valor: 8000, Quantia: 2},
	}
	orcamento := &orcamento.Orcamento{
		Produtos: produtos,
	}

	fmt.Println(">> must be", calculadoraPIS.RealizaCalculo(orcamento))
}
