package orcamento

type Produto struct {
	Nome    string
	Valor   float64
	Quantia int
}

type Orcamento struct {
	Produtos       []Produto
	ValorImpostos  float64
	ValorDescontos float64
}

func (o *Orcamento) GetValorOrcamento() float64 {
	valorProdutos := 0.0
	for _, produto := range o.Produtos {
		valorProdutos += float64(produto.Quantia) * produto.Valor
	}
	return o.ValorImpostos + valorProdutos
}
