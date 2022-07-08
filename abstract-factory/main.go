package main

import "fmt"

func main() {
	nikeFactory := getApparelFactory("nike")
	adidasFactory := getApparelFactory("adidas")

	nikeShoe := nikeFactory.CreateShoe()
	adidasShoe := adidasFactory.CreateShoe()

	fmt.Println(nikeShoe.GetLogo())
	fmt.Println(adidasShoe.GetLogo())
}

type iShoe interface {
	SetLogo(logo string)
	GetLogo() string
}

type shoe struct {
	logo string
}

type BrandAbstractFactory interface {
	CreateShoe() iShoe
}

type NikeFactory struct{}

type NikeShoe struct {
	shoe
}

func (n *NikeShoe) SetLogo(logo string) {
	n.logo = logo
}

func (n *NikeShoe) GetLogo() string {
	return n.logo
}

func (nike *NikeFactory) CreateShoe() iShoe {
	return &NikeShoe{
		shoe{
			logo: "Nike",
		},
	}
}

type AdidasFactory struct {
}

type AdidasShoe struct {
	shoe
}

func (n *AdidasShoe) SetLogo(logo string) {
	n.logo = logo
}

func (n *AdidasShoe) GetLogo() string {
	return n.logo
}

func (nike *AdidasFactory) CreateShoe() iShoe {
	return &AdidasShoe{
		shoe{
			logo: "Adidas",
		},
	}
}
func getApparelFactory(brand string) BrandAbstractFactory {
	switch brand {
	case "nike":
		return &NikeFactory{}
	case "adidas":
		return &AdidasFactory{}
	default:
		return &NikeFactory{}
	}
}
