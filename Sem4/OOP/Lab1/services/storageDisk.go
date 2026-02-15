package services

import "errors"

const (
	HDD = "hdd"
	SSD = "ssd"
)

type Disk struct {
	diskType    string
	capacity    uint64
	isConnected bool
	rwSpeed     uint64
}

func NewDisk(diskType string, capacity uint64, isConnected bool) *Disk {
	return &Disk{
		diskType:    diskType,
		capacity:    capacity,
		isConnected: isConnected,
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
