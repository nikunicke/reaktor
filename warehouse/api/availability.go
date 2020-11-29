package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/reaktor/warehouse"
)

var _ warehouse.AvailabilityService = &AvailabilityService{}

type AvailabilityService struct {
	c *Client
}

func NewAvailabilityService(c *Client) *AvailabilityService {
	return &AvailabilityService{c: c}
}

func (s *AvailabilityService) GetAvailability(
	manufacturer string) (*warehouse.Availability, error) {
	prefix := "availability/"
	resp, err := s.c.Get(prefix + manufacturer)
	if err != nil {
		return nil, err
	}
	return unmarshalAvailabilityResponse(resp)
}

func unmarshalAvailabilityResponse(
	resp *http.Response) (*warehouse.Availability, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var availability warehouse.Availability
	if err := json.Unmarshal(body, &availability); err != nil {
		fmt.Println("error here?")
		return nil, err
	}
	return &availability, nil
}
