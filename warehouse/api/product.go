package api

import (
	"encoding/json"
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
	prefix := "products/"
	resp, err := s.c.Get(prefix + t)
	if err != nil {
		return nil, err
	}
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
