package warehouse_api

import "encoding/xml"

type AvailabilityService interface {
	GetAvailability(manufacturer string) (*Availability, error)
}

type Availability struct {
	Code     int        `json:"code"`
	Response []Response `json:"response"`
}

type Response struct {
	ID          string `json:"id"`
	DataPayload string `json:"DATAPAYLOAD"`
}

type AvailabilityXML struct {
	Availability xml.Name `xml:"AVAILABILITY"`
	InStockValue string   `xml:"INSTOCKVALUE"`
}
