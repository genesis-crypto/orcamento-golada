package orcamento

import "fmt"

type EmAprovacao struct {
	orcamento *Orcamento
}

func (e *EmAprovacao) AdicionaDescontoExtra() error {
	fmt.Println("Info(desconto extra): State EmAprovação desconto aplicado")
	e.orcamento.ValorDescontos += e.orcamento.GetValorOrcamento() * 0.05
	return nil
}

func (e *EmAprovacao) Aprova() error {
	fmt.Println("\nInfo(transição): State em transição EmAprovação -> Aprovado")
	e.orcamento.SetState(e.orcamento.aprovado)
	return nil
}

func (e *EmAprovacao) Reprova() error {
	fmt.Println("\nInfo(transição): State em transição EmAprovação -> Aprovado")
	e.orcamento.SetState(e.orcamento.reprovado)
	return nil
}

func (e *EmAprovacao) Finaliza() error {
	return fmt.Errorf("Error(state): Orcamento em aprovação não podem ir para finalizado diretamente")
}
