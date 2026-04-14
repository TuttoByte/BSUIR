package sorts

import (
	"Lab4/services/models/rerq_types"
	"sort"
)

type MassSort struct {
}

func (m *MassSort) Sort(data []rerq_types.DeliveryResponse) {
	sort.Slice(data, func(i, j int) bool {
		if data[i].Time > data[j].Time {
			return true
		}
		return false
	})
}
