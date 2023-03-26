package desconto

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type DescontoPorQuantia struct {
	next IHDesconto
}

func (d *DescontoPorQuantia) Executa(orcamento *orcamento.Orcamento) {
	totalQuantiaProdutos := 0.0
	for _, produto := range orcamento.Produtos {
		totalQuantiaProdutos += float64(produto.Quantia)
	}

	if totalQuantiaProdutos < 5.0 {
		fmt.Println("Info(desconto): A quantia de produtos deve ser maior que 5 unidades")
		return
	}
	fmt.Println("Info(desconto): Desconto aplicado - DescontoPorQuantia")
	orcamento.ValorDescontos += orcamento.GetValorOrcamento() * 0.1
	fmt.Println("Info(desconto): Desconto total aplicado R$", orcamento.ValorDescontos)
	d.next.Executa(orcamento)
}

func (d *DescontoPorQuantia) SetNext(next IHDesconto) {
	d.next = next
}
