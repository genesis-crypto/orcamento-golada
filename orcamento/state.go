package orcamento

type State interface {
	AdicionaDescontoExtra() error
	Aprova() error
	Reprova() error
	Finaliza() error
}
