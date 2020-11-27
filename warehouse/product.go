package warehouse

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
	Availability string   `json:"availability"`
}

// ProductCategories list all product types
var ProductCategories = [...]string{Jackets, Shirts, Accessories}

// Defined product categoris
const (
	Jackets     = string("products/jackets")
	Shirts      = string("products/shirts")
	Accessories = string("products/accessories")
)

// Product errors
const (
	ErrorInvalidCategory = Error("Invalid category")
)

// IsValidProductCategory checks if category exists
func IsValidProductCategory(ctg string) error {
	switch ctg {
	case Jackets, Shirts, Accessories:
		return nil
	}
	return ErrorInvalidCategory
}
