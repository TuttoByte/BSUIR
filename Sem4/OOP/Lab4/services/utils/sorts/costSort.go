package sorts

import (
	"Lab4/services/models/rerq_types"
	"sort"
)

type CostSort struct {
}

func (c *CostSort) Sort(data []rerq_types.DeliveryResponse) {
	sort.Slice(data, func(i, j int) bool {
		if data[i].TotoalCost >= data[j].TotoalCost {
			return true
		}
		return false
	})
}
