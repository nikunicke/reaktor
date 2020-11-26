package warehouse_api

import "encoding/xml"

type AvailabilityService interface {
	GetAvailability(manufacturer string) (Availabilities, error)
}

type Availabilities []Availability

type Availability struct {
	ID          string
	DataPayload []byte
}

type AvailabilityXML struct {
	Availability xml.Name `xml:"AVAILABILITY"`
	InStockValue string   `xml:"INSTOCKVALUE"`
}
