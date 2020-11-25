package warehouse_api

import "encoding/xml"

type Availability struct {
	ID          string
	DataPayload []byte
}

type AvailabilityXML struct {
	Availability xml.Name `xml:"AVAILABILITY"`
	InStockValue string   `xml:"INSTOCKVALUE"`
}
