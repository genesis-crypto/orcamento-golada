package orcamento

import "fmt"

type Finalizado struct {
	orcamento *Orcamento
}

func (r *Finalizado) AdicionaDescontoExtra() error {
	return fmt.Errorf("Error(state): Orçamento não pode receber descontos extras em estado de finalizado")
}

func (a *Finalizado) Aprova() error {
	return fmt.Errorf("Error(state): Orçamento está em estado de finalizado e não pode ser aprovado")
}

func (a *Finalizado) Reprova() error {
	return fmt.Errorf("Error(state): Orçamento está em estado de finalizado e não pode ser reprovado")
}

func (a *Finalizado) Finaliza() error {
	return fmt.Errorf("Error(state): Orçamento já está em estado de finalizado")
}
