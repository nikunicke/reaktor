package warehouse

import (
	"fmt"
	"strings"
	"time"
)

type WarehouseService interface {
	GetProducts(category string) (Products, error)
}

type WarehouseWorker interface {
	Update() error
}

type Warehouse struct {
	Inventory           Inventory
	UpdatedAt           time.Time
	UpdateInterval      *time.Ticker
	ProductService      ProductService
	AvailabilityService AvailabilityService
}

type Inventory map[string]Products

func New(categories ...string) *Warehouse {
	c := make(Inventory)
	for _, category := range categories {
		if _, ok := c[category]; !ok {
			c[category] = Products{}
		}
	}
	return &Warehouse{
		Inventory: c,
	}
}

func (w *Warehouse) Update() error {
	p := make(chan Products)
	a := make(chan AvailabilityResponseMap)
	e := make(chan error)
	inventory := w.Inventory

	fmt.Println("Updating warehouse data")
	for _, category := range ProductCategories {
		go func(ctg string) {
			products, err := w.ProductService.GetProducts(ctg)
			if err != nil {
				fmt.Println(err)
				e <- err
			}
			for range Manufacturers {
				select {
				case availability := <-a:
					{
						for i, _ := range products {
							key := strings.ToUpper(products[i].ID)
							if _, ok := availability[key]; ok {
								products[i].Availability = availability[key]
							}
						}
					}
				}
			}
			p <- products
		}(category)
	}

	for m := range Manufacturers {
		go func(manufacturer string) {
			availability, err := w.AvailabilityService.GetAvailability(manufacturer)
			if err != nil {
				fmt.Println(err)
				for range ProductCategories {
					e <- err
				}
			}
			for range ProductCategories {
				a <- availability.Response.Map()
			}
		}(m)
	}

	for range ProductCategories {
		select {
		case err := <-e:
			return err
		case productData := <-p:
			inventory[productData[0].Type] = productData
		}
	}
	w.Inventory = inventory
	w.UpdatedAt = time.Now()
	fmt.Println("Update complete")
	return nil
}

func (w *Warehouse) GetProducts(category string) (Products, error) {
	if _, ok := w.Inventory[category]; ok {
		return w.Inventory[category], nil
	}
	return nil, ErrorInvalidCategory
}
