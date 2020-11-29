package warehouse

type ProductService interface {
	GetProducts(t string) (Products, error)
}

// Product errors
const (
	ErrorInvalidCategory = Error("Invalid category")
)

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
var ProductCategories = []string{Jackets, Shirts, Accessories}

// Defined product categoris
const (
	Jackets     = string("jackets")
	Shirts      = string("shirts")
	Accessories = string("accessories")
)

// Pre defined manufacturers.
const (
	Derp    = string("derp")
	Reps    = string("reps")
	Xoon    = string("xoon")
	Nouke   = string("nouke")
	Abiplos = string("abiplos")
)

// Manufacturers map can be used as is, but if additional manufacturers
// should be expected, either ask the business owner and manually add items
// in here, or iterate trough Products to find them, and then add them in here
var Manufacturers = map[string]struct{}{
	Derp:    struct{}{},
	Reps:    struct{}{},
	Xoon:    struct{}{},
	Nouke:   struct{}{},
	Abiplos: struct{}{},
}

func (p *Products) GetManufacturers(m map[string]struct{}) map[string]struct{} {
	for _, item := range *p {
		if _, ok := m[item.Manufacturer]; !ok {
			m[item.Manufacturer] = struct{}{}
		}
	}
	return m
}

// IsValidProductCategory checks if category exists
func IsValidProductCategory(ctg string) error {
	switch ctg {
	case Jackets, Shirts, Accessories:
		return nil
	}
	return ErrorInvalidCategory
}
