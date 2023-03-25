package calculadora

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type CalculadoraPIS struct {
	Calculadora ICalculadora
}

func (c *CalculadoraPIS) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	orcamento.ValorImpostos = orcamento.ValorImpostos + 1
	fmt.Println("> PIS ", orcamento.ValorImpostos)

	return c.Calculadora.RealizaCalculo(orcamento)
}
