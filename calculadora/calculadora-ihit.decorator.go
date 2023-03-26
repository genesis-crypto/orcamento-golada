package calculadora

import (
	"fmt"
	"orcamento-golada/orcamento"
)

type CalculadoraIHIT struct {
	Calculadora ICalculadora
}

func (c *CalculadoraIHIT) RealizaCalculo(orcamento *orcamento.Orcamento) float64 {
	temItensComMesmoNome := false
	nomesProdutos := make(map[string]int, 2)
	for _, produto := range orcamento.Produtos {
		nomesProdutos[produto.Nome]++
		if nomesProdutos[produto.Nome] > 1 {
			temItensComMesmoNome = true
		}
	}

	if temItensComMesmoNome {
		valorImpostoIHIT := (orcamento.GetValorOrcamento() * 0.13) + 100
		orcamento.ValorImpostos = orcamento.ValorImpostos + valorImpostoIHIT
	} else {
		valorImpostoIHIT := 0.1 * float64(len(orcamento.Produtos))
		orcamento.ValorImpostos = orcamento.ValorImpostos + valorImpostoIHIT
	}

	fmt.Println("> IHIT ", orcamento.ValorImpostos)

	return c.Calculadora.RealizaCalculo(orcamento)
}
