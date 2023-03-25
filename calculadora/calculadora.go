package calculadora

import "orcamento-golada/orcamento"

type Calculadora struct{}

func (c *Calculadora) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	return orcamento.GetValorOrcamento()
}
