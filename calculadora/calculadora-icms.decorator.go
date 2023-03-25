package calculadora

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type CalculadoraICMS struct {
	Calculadora ICalculadora
}

func (c *CalculadoraICMS) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	orcamento.ValorImpostos = orcamento.ValorImpostos + 1
	fmt.Println("> ICMS ", orcamento.ValorImpostos)

	return c.Calculadora.RealizaCalculo(orcamento)
}
