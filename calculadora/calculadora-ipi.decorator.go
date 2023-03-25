package calculadora

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type CalculadoraIPI struct {
	Calculadora ICalculadora
}

func (c *CalculadoraIPI) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	orcamento.ValorImpostos = orcamento.ValorImpostos + 1
	fmt.Println("> IPI ", orcamento.ValorImpostos)

	return c.Calculadora.RealizaCalculo(orcamento)
}
