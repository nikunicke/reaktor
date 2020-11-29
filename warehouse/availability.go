package warehouse

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

type AvailabilityResponseMap map[string]string

func (r AvailabilityResponse) Map() AvailabilityResponseMap {
	availability := make(AvailabilityResponseMap)
	for _, item := range r {
		availability[item.ID] = item.DataPayload
	}
	return availability
}
