package online_shopping

type ProductCategory int

const (
	CLOTH ProductCategory = iota
	SHOES
	ELECTRONICS
)

type Product struct {
	Id          string
	Name        string
	Description string
	Category    ProductCategory
	Count       int
	Price       float64
}

func NewProduct(id, name, description string, category ProductCategory, initialAvailability int, price float64) *Product {
	return &Product{
		Id:          id,
		Name:        name,
		Description: description,
		Category:    category,
		Count:       initialAvailability,
		Price:       price,
	}
}

func (p *Product) IsAvailable() bool {
	return p.Count > 0
}

func (p *Product) UpdateQuantity(quantity int) {
	p.Count += quantity
}
