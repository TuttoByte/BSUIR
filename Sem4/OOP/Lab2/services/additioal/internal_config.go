package additioal

import "net"

type DeviceInfo struct {
	MaxTemp float32
	uapi    net.Listener
}

func (d *DeviceInfo) GetMaxTemp() float32 {
	return d.MaxTemp
}

func (d *DeviceInfo) GetUAPI() net.Listener {
	return d.uapi
}
