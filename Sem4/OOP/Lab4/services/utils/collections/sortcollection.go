package collections

import (
	"Lab4/services/models/rerq_types"
	"Lab4/services/utils/contracts"
)

type SortCollection struct {
	data  []*rerq_types.DeliveryResponse
	index int
}

func NewInterator() *SortCollection {
	return &SortCollection{
		data: make([]*rerq_types.DeliveryResponse, 0),
	}
}

func (s *SortCollection) HasNext() bool {
	if s.index < len(s.data)-1 {
		return true
	}
	return false
}
func (s *SortCollection) Next() *rerq_types.DeliveryResponse {
	if s.HasNext() {
		elem := s.data[s.index]
		s.index++
		return elem
	}
	return nil
}

func (s *SortCollection) Append(data []*rerq_types.DeliveryResponse) {
	s.data = append(s.data, data...)
}
func (s *SortCollection) Len() int {
	return len(s.data)
}

func (s *SortCollection) PushBask(elem *rerq_types.DeliveryResponse) {
	s.data = append(s.data, elem)
}

func (s *SortCollection) Sort(sm contracts.Sortable) {
	sm.Sort(s.data)
}
