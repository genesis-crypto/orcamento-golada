package nota_fiscal

import "fmt"

type Client struct {
	Id string
}

func (c *Client) Update(itemName string) {
	fmt.Printf("\nEnviando para email %s nf de cnpj %s\n", c.Id, itemName)
}

func (c *Client) GetID() string {
	return c.Id
}
