package contracts

import "Lab4/services/models/rerq_types"

type Sortable interface {
	Sort([]rerq_types.DeliveryResponse)
}
