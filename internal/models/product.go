package models

const ProductValidationNameRequired = "The Name field is required."
const ProductValidationPriceGte = "The Price must be greater than or equal 0."

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p Product) Validate() []string {
	messages := []string{}
	if p.Name == "" {
		messages = append(messages, ProductValidationNameRequired)
	}
	if p.Price < 0 {
		messages = append(messages, ProductValidationPriceGte)
	}
	return messages
}

func (p *Product) Fill(params struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}) {
	p.Name = params.Name
	p.Price = params.Price
}
