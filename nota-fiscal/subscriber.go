package nota_fiscal

type Subscriber interface {
	Update(string)
	GetID() string
}
