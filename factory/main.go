package main

import (
	"fmt"

	"github.com/rikutkb/go_design_pattern/factory/framework"
)

type IDCard struct {
	owner string
}

func (card *IDCard) Use() {
	fmt.Println(card.owner + "のカードを使用します。")
}
func (card *IDCard) getOwner() string {
	return card.owner
}

type IDCardFactory struct {
	owners []framework.IProduct
}

func (f *IDCardFactory) CreateProduct(owner string) framework.IProduct {
	return &IDCard{owner: owner}
}

func (f *IDCardFactory) RegisterProduct(product framework.IProduct) {
	f.owners = append(f.owners, product)
}
func (f *IDCardFactory) getOwners() []framework.IProduct {
	return f.owners
}

func main() {
	factory := framework.Factory{IFactory: &IDCardFactory{}}
	card1 := factory.IFactory.CreateProduct("1")
	card2 := factory.IFactory.CreateProduct("2")
	card1.Use()
	card2.Use()
}
