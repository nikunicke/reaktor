package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/reaktor/warehouse"
)

var _ warehouse.ProductService = &ProductService{}

type ProductService struct {
	c *Client
}

func NewProductService(c *Client) *ProductService {
	return &ProductService{c: c}
}

func (s *ProductService) GetProducts(
	t string) (warehouse.Products, error) {
	resp, err := s.c.Get(t)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return unmarshalProductResponse(resp)
}

func unmarshalProductResponse(
	resp *http.Response) (warehouse.Products, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var products warehouse.Products
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}
	return products, nil
}
