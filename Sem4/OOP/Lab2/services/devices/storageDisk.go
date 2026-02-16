package devices

import (
	"Lab1/services/additioal"
	"errors"
	"net"
)

const (
	HDD = "hdd"
	SSD = "ssd"
)

type Disk struct {
	diskType    string
	capacity    uint64
	isConnected bool
	rwSpeed     uint64
	info        additioal.DeviceInfo
}

func NewDisk(diskType string, capacity uint64) *Disk {
	return &Disk{
		diskType:    diskType,
		capacity:    capacity,
		isConnected: true,
	}
}

func (d *Disk) DiskType() string {
	return d.diskType
}

func (d *Disk) Connect() error {
	if d.isConnected {
		return errors.New("already started")
	}
	d.isConnected = true
	return nil
}

func (d *Disk) Disconnect() error {
	if !d.isConnected {
		return errors.New("already disconnected")
	}
	d.isConnected = false
	return nil
}
func (d *Disk) GetUAPI() net.Listener {
	return d.info.GetUAPI()
}
func (d *Disk) GetMaxTemp() float32 {
	return d.info.GetMaxTemp()
}
func (d *Disk) GetType() string {
	return "Disk"
}
