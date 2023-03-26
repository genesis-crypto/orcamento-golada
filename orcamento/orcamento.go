package orcamento

type Produto struct {
	Nome    string
	Valor   float64
	Quantia int
}

type Orcamento struct {
	emAprovacao State
	aprovado    State
	reprovado   State
	finalizado  State
	stateAtual  State

	Produtos       []Produto
	ValorImpostos  float64
	ValorDescontos float64
}

func CriaOrcamento(produtos []Produto) *Orcamento {
	o := &Orcamento{
		Produtos:       produtos,
		ValorImpostos:  0.0,
		ValorDescontos: 0.0,
	}

	emAprovacao := &EmAprovacao{
		orcamento: o,
	}

	aprovado := &Aprovado{
		orcamento: o,
	}

	reprovado := &Reprovado{
		orcamento: o,
	}

	finalizado := &Finalizado{
		orcamento: o,
	}

	o.SetState(emAprovacao)
	o.emAprovacao = emAprovacao
	o.aprovado = aprovado
	o.reprovado = reprovado
	o.finalizado = finalizado

	return o
}

func (o *Orcamento) GetValorOrcamento() float64 {
	valorProdutos := 0.0
	for _, produto := range o.Produtos {
		valorProdutos += float64(produto.Quantia) * produto.Valor
	}
	return o.ValorImpostos + valorProdutos
}

func (o *Orcamento) AdicionaDescontoExtra() error {
	return o.stateAtual.AdicionaDescontoExtra()
}

func (o *Orcamento) Aprova() error {
	return o.stateAtual.Aprova()
}

func (o *Orcamento) Reprova() error {
	return o.stateAtual.Reprova()
}

func (o *Orcamento) Finaliza() error {
	return o.stateAtual.Finaliza()
}

func (o *Orcamento) SetState(s State) {
	o.stateAtual = s
}
