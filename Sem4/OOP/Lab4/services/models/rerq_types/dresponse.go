package rerq_types

type DeliveryResponse struct {
	Type        string  `json:"type" csv:"type"`
	Mass        float64 `json:"mass" csv:"mass"`
	Destination float64 `json:"destination" csv:"destination"`
	TotoalCost  float64 `json:"totoal_cost" csv:"totoal_cost"`
}
