package services

import (
	"errors"
	"fmt"
)

type RAM struct {
	capacity  uint64
	frequency uint64
	isStarted bool
}

func NewRAM(capacity uint64, frequency uint64) *RAM {
	return &RAM{
		capacity:  capacity,
		frequency: frequency,
	}
}

func (r *RAM) Capacity() uint64 {
	return r.capacity
}

func (r *RAM) Frequency() uint64 {
	return r.frequency
}

func (r *RAM) IsStarted() bool {
	return r.isStarted
}

func (r *RAM) Start() {
	if r.isStarted {
		err := errors.New("RAM is already started")
		fmt.Println(err)
	}
	r.isStarted = true
}
