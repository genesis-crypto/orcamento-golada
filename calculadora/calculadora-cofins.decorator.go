package calculadora

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type CalculadoraCofins struct {
	Calculadora ICalculadora
}

func (c *CalculadoraCofins) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	orcamento.ValorImpostos = orcamento.ValorImpostos + 1
	fmt.Println("> Cofins ", orcamento.ValorImpostos)

	return c.Calculadora.RealizaCalculo(orcamento)
}
