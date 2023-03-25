package calculadora

type CalculadoraICMS struct {
	Calculadora ICalculadora
}

func (c *CalculadoraICMS) RealizaCalculo(orcamento float64) float64 {
	orcamento = orcamento + 2.1
	return c.Calculadora.RealizaCalculo(orcamento)
}
