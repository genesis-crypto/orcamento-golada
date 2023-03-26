package desconto

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type DescontoPorValor struct {
	next IHDesconto
}

func (d *DescontoPorValor) Executa(orcamento *orcamento.Orcamento) {
	totalValorOrcamento := 0.0
	for _, produto := range orcamento.Produtos {
		totalValorOrcamento += produto.Valor
	}

	if totalValorOrcamento < 500.0 {
		fmt.Println("Info(desconto): O valor do orçamento deve ser maior que R$ 500,00")
		return
	}

	fmt.Println("Info(desconto): Desconto aplicado - DescontoPorValor")
	orcamento.ValorDescontos += orcamento.GetValorOrcamento() * 0.07
	fmt.Println("Info(desconto): Desconto total aplicado R$", orcamento.ValorDescontos)
}

func (d *DescontoPorValor) SetNext(next IHDesconto) {
	d.next = next
}
