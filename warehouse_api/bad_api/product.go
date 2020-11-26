package bad_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/reaktor/warehouse_api"
)

var _ warehouse_api.ProductService = &ProductService{}

type ProductService struct {
	c *Client
}

func NewProductService(c *Client) *ProductService {
	return &ProductService{c: c}
}

func (s *ProductService) GetProducts(
	t string) (warehouse_api.Products, error) {
	fmt.Println(t)
	resp, err := s.c.Get(t)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return unmarshalResponse(resp)
}

func unmarshalResponse(
	resp *http.Response) (warehouse_api.Products, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var products warehouse_api.Products
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}
	return products, nil
}
