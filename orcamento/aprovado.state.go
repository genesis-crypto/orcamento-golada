package orcamento

import "fmt"

type Aprovado struct {
	orcamento *Orcamento
}

func (a *Aprovado) AdicionaDescontoExtra() error {
	fmt.Println("Info(desconto extra): State Aprovado desconto aplicado")
	a.orcamento.ValorDescontos += a.orcamento.GetValorOrcamento() * 0.02
	return nil
}

func (a *Aprovado) Aprova() error {
	return fmt.Errorf("Error(state): Orçamento já está em estado de aprovação")
}

func (a *Aprovado) Reprova() error {
	return fmt.Errorf("Error(state): Orçamento está em estado de aprovação e não pode ser reprovado")
}

func (a *Aprovado) Finaliza() error {
	fmt.Println("\nInfo(transição): State em transição Aprovado -> Finalizado")
	a.orcamento.SetState(a.orcamento.finalizado)
	return nil
}
