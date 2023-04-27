package nota_fiscal

import "time"

type Publisher interface {
	AddSubscriber(subscriber Subscriber)
	RemoveSubscriber(subscriber Subscriber)
	notifySubscribers()
}

type NotaFiscal struct {
	razaoSocial string
	cnpj        string
	valorTotal  float64
	impostos    float64
	data        time.Time
	observacoes string
}

type NotaFiscalBuilder struct {
	razaoSocial string
	cnpj        string
	valorTotal  float64
	impostos    float64
	data        time.Time
	observacoes string
	subscribers []Subscriber
}

func (b *NotaFiscalBuilder) SetRazaoSocial(razaoSocial string) *NotaFiscalBuilder {
	b.razaoSocial = razaoSocial
	return b
}

func (b *NotaFiscalBuilder) SetCnpj(cnpj string) *NotaFiscalBuilder {
	b.cnpj = cnpj
	return b
}

func (b *NotaFiscalBuilder) SetValorTotal(valorTotal float64) *NotaFiscalBuilder {
	b.valorTotal = valorTotal
	return b
}

func (b *NotaFiscalBuilder) SetImpostos(impostos float64) *NotaFiscalBuilder {
	b.impostos = impostos
	return b
}

func (b *NotaFiscalBuilder) SetData(data time.Time) *NotaFiscalBuilder {
	b.data = data
	return b
}

func (b *NotaFiscalBuilder) SetObservacoes(observacoes string) *NotaFiscalBuilder {
	b.observacoes = observacoes
	return b
}

func (b *NotaFiscalBuilder) Build() NotaFiscal {
	nf := NotaFiscal{
		razaoSocial: b.razaoSocial,
		cnpj:        b.cnpj,
		valorTotal:  b.valorTotal,
		impostos:    b.impostos,
		data:        b.data,
		observacoes: b.observacoes,
	}
	b.notifySubscribers()
	return nf
}

func (n *NotaFiscalBuilder) AddSubscriber(subscriber Subscriber) {
	n.subscribers = append(n.subscribers, subscriber)
}

func (n *NotaFiscalBuilder) RemoveSubscriber(subscriber Subscriber) {
	for i, obs := range n.subscribers {
		if obs == subscriber {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NotaFiscalBuilder) notifySubscribers() {
	for _, subscriber := range n.subscribers {
		subscriber.Update(n.cnpj)
	}
}
