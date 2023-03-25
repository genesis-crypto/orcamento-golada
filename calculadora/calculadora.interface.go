package calculadora

import "orcamento-golada/orcamento"

type ICalculadora interface {
	RealizaCalculo(orcamento *orcamento.Orcamento) float64
}
