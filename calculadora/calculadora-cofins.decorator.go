package calculadora

type CalculadoraCofins struct {
	Calculadora ICalculadora
}

func (c *CalculadoraCofins) RealizaCalculo(orcamento float64) float64 {
	orcamento = orcamento + 1
	return c.Calculadora.RealizaCalculo(orcamento)
}
