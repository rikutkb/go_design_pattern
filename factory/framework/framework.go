package framework

type IProduct interface {
	Use()
}

type IFactory interface {
	CreateProduct(owner string) IProduct
	RegisterProduct(product IProduct)
}

type Factory struct {
	IFactory IFactory
}

func (f *Factory) Create(owner string) IProduct {
	p := f.IFactory.CreateProduct(owner)
	f.IFactory.RegisterProduct(p)
	return p
}
