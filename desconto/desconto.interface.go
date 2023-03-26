package desconto

import "orcamento-golada/orcamento"

type IHDesconto interface {
	SetNext(h IHDesconto)
	Executa(orcamento *orcamento.Orcamento)
}
