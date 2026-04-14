package rerq_types

type DeliveryRequest struct {
	Batches []Batch `json:"batches"`
}

type Batch struct {
	CargoNumber       int    `json:"cargo_number" xml:"cargo_number"`
	CargoType         string `json:"cargo_type" xml:"cargo_type"`
	DeliveryDistance  int    `json:"delivery_distance" xml:"delivery_distance"`
	TransportType     string `json:"transport_type" xml:"transport_type"`
	TransportInstance string `json:"transport_instance" xml:"transport_instance"`
}
