package devices

import (
	"Lab1/services/additioal"
	"errors"
	"net"
)

type RAM struct {
	capacity  uint64
	frequency uint64
	isStarted bool
	info      additioal.DeviceInfo
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

func (r *RAM) Connect() error {
	if r.isStarted {
		err := errors.New("RAM is already started")
		return err
	}
	r.isStarted = true
	return nil
}

func (r *RAM) Disconnect() error {
	r.isStarted = false
	return nil
}

func (r *RAM) GetType() string {
	return "RAM"
}

func (r *RAM) GetUAPI() net.Listener {
	return r.info.GetUAPI()
}
func (r *RAM) GetMaxTemp() float32 {
	return r.info.GetMaxTemp()
}
