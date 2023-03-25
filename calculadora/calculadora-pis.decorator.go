package calculadora

type CalculadoraPIS struct {
	Calculadora ICalculadora
}

func (c *CalculadoraPIS) RealizaCalculo(orcamento float64) float64 {
	orcamento = orcamento + 1.0
	return c.Calculadora.RealizaCalculo(orcamento)
}
