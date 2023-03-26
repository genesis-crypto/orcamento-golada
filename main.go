package main

import (
	"fmt"
	"log"
	"orcamento-golada/calculadora"
	"orcamento-golada/desconto"
	"orcamento-golada/orcamento"
)

func main() {
	calculadoraBase := &calculadora.Calculadora{}

	calculadoraICMS := &calculadora.CalculadoraICMS{
		Calculadora: calculadoraBase,
	}

	calculadoraPIS := &calculadora.CalculadoraPIS{
		Calculadora: calculadoraICMS,
	}

	orcamentoPrincipal := orcamento.CriaOrcamento([]orcamento.Produto{
		{Nome: "Notebook", Valor: 8000, Quantia: 10},
	})

	fmt.Println(">> must be", calculadoraPIS.RealizaCalculo(orcamentoPrincipal))

	descontoValor := &desconto.DescontoPorValor{}

	descontoQuantia := &desconto.DescontoPorQuantia{}
	descontoQuantia.SetNext(descontoValor)

	descontoQuantia.Executa(orcamentoPrincipal)

	orcamentoPrincipal.AdicionaDescontoExtra()

	if err := orcamentoPrincipal.Aprova(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := orcamentoPrincipal.Finaliza(); err != nil {
		log.Fatalf(err.Error())
	}
}
