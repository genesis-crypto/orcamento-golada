package main

import (
	"fmt"
	"orcamento-golada/calculadora"
	"orcamento-golada/desconto"
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
		{Nome: "Notebook", Valor: 8000, Quantia: 10},
	}

	orcamento := &orcamento.Orcamento{
		Produtos: produtos,
	}

	fmt.Println(">> must be", calculadoraPIS.RealizaCalculo(orcamento))

	descontoValor := &desconto.DescontoPorValor{}

	descontoQuantia := &desconto.DescontoPorQuantia{}
	descontoQuantia.SetNext(descontoValor)

	descontoQuantia.Executa(orcamento)
}
