package bad_api

import "github.com/nikunicke/reaktor/warehouse_api"

var _ warehouse_api.AvailabilityService = &AvailabilityService{}

type AvailabilityService struct {
	c *Client
}

func NewAvailabilityService(c *Client) *AvailabilityService {
	return &AvailabilityService{c: c}
}

func (s *AvailabilityService) GetAvailability(
	manufacturer string) (warehouse_api.Availabilities, error) {
	return nil, nil
}
