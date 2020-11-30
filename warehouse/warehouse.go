package warehouse

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type WarehouseService interface {
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
	a := make(chan []*AvailabilityResponseMap)
	quit := make(chan bool)
	done := make(chan error)
	Availabilities := make([]*AvailabilityResponseMap, len(Manufacturers))
	Inventory := make(Inventory)
	var lock = sync.RWMutex{}

	fmt.Println("Updating warehouse data")
	for _, category := range ProductCategories {
		go func(category string) {
			var err error
			var products Products
			if products, err = w.ProductService.GetProducts(category); err == nil {
				availabilities := <-a
				for _, availability := range availabilities {
					for i, _ := range products {
						key := strings.ToUpper(products[i].ID)
						if _, ok := (*availability)[key]; ok {
							products[i].Availability = (*availability)[key]
						}
					}
				}
				lock.Lock()
				defer lock.Unlock()
				Inventory[category] = products
			}
			select {
			case done <- err:
			case <-quit:
			}
		}(category)
	}

	var i = 0
	for m := range Manufacturers {
		go func(i int, manufacturer string) {
			var value AvailabilityResponseMap
			var availability *Availability
			var err error
			if availability, err = w.AvailabilityService.GetAvailability(manufacturer); err == nil {
				value, err = availability.Response.Map()
				Availabilities[i] = &value
			}
			select {
			case done <- err:
			case <-quit:
			}
		}(i, m)
		i++
	}
	for success := 0; success < 5; success++ {
		err := <-done
		if err != nil {
			close(quit)
			fmt.Println("Updated failed")
			return err
		}
	}
	a <- Availabilities
	a <- Availabilities
	a <- Availabilities
	for success := 0; success < 3; success++ {
		err := <-done
		if err != nil {
			close(quit)
			fmt.Println("Updated failed")
			return err
		}
	}
	w.Inventory = Inventory
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
