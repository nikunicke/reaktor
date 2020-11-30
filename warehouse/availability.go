package warehouse

import (
	"encoding/xml"
)

type AvailabilityService interface {
	GetAvailability(manufacturer string) (*Availability, error)
}

type Availability struct {
	Code     int                   `json:"code"`
	Response *AvailabilityResponse `json:"response"`
}

type Response struct {
	ID          string `json:"id"`
	DataPayload string `json:"DATAPAYLOAD"`
}

type AvailabilityResponse []Response

type AvailabilityXML struct {
	XMLName      xml.Name `xml:"AVAILABILITY"`
	InStockValue string   `xml:"INSTOCKVALUE"`
}

type AvailabilityResponseMap map[string]string

func (r AvailabilityResponse) Map() (AvailabilityResponseMap, error) {
	availability := make(AvailabilityResponseMap)
	var xmlData AvailabilityXML
	for _, item := range r {
		if err := xml.Unmarshal([]byte(item.DataPayload), &xmlData); err != nil {
			return nil, err
		}
		availability[item.ID] = xmlData.InStockValue
	}
	return availability, nil
}
