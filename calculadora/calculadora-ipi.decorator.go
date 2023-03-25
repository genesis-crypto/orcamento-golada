package calculadora

type CalculadoraIPI struct {
	Calculadora ICalculadora
}

func (c *CalculadoraIPI) RealizaCalculo(orcamento float64) float64 {
	orcamento = orcamento + 1
	return c.Calculadora.RealizaCalculo(orcamento)
}
