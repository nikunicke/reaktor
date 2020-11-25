package warehouse_api

type ProductService interface {
	GetProducts(t ProductType) (Products, error)
}

type Products []Product

type Product struct {
	ID           string      `json:"id"`
	Type         ProductType `json:"type"`
	Name         string      `json:"name"`
	Color        []string    `json:"color"`
	Price        int         `json:"price"`
	Manufacturer string      `json:"manufacturer"`
}

type ProductType string

const (
	Jacket      = ProductType("jackets")
	Shirt       = ProductType("shirts")
	Accessories = ProductType("accessories")
)

func (t ProductType) String() string { return string(t) }
