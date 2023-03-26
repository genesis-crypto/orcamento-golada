package orcamento

import "fmt"

type Reprovado struct {
	orcamento *Orcamento
}

func (r *Reprovado) AdicionaDescontoExtra() error {
	return fmt.Errorf("Error(state): Orçamento não pode receber descontos extras em estado de Reprovado")
}

func (a *Reprovado) Aprova() error {
	return fmt.Errorf("Error(state): Orçamento está em estado de reprovado e não pode ser aprovado")
}

func (a *Reprovado) Reprova() error {
	return fmt.Errorf("Error(state): Orçamento já está em estado de reprovado")
}

func (a *Reprovado) Finaliza() error {
	fmt.Println("\nInfo(transição): State em transição Reprovado -> Finalizado")
	a.orcamento.SetState(a.orcamento.finalizado)
	return nil
}
