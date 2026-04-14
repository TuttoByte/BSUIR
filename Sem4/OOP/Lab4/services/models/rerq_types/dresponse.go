package rerq_types

type DeliveryResponse struct {
	Type        string  `json:"type" csv:"type" xml:"type"`
	Time        float64 `json:"time" csv:"time" xml:"time"`
	Destination float64 `json:"destination" csv:"destination" xml:"destination"`
	TotoalCost  float64 `json:"totoal_cost" csv:"totoal_cost" xml:"totoal_cost"`
}
