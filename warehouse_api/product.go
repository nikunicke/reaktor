package warehouse_api

type ProductService interface {
	GetProducts(t string) (Products, error)
}

type Products []Product

type Product struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	Color        []string `json:"color"`
	Price        int      `json:"price"`
	Manufacturer string   `json:"manufacturer"`
}

// type ProductType string

const (
	Jackets     = string("products/jackets")
	Shirts      = string("products/shirts")
	Accessories = string("products/accessories")
)

// func (t ProductType) String() string { return string(t) }
