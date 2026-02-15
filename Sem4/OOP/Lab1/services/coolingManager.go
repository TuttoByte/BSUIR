package services

import (
	"errors"
	"time"
)

type MaxCoolingTemp struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Name    string
	MaxTemp float32
}

type CoolingManager struct {
	devices      map[uint64]float32
	MaxTemp      float32
	isMonitoring bool
}

func NewCoolingManager(maxTemp float32) *CoolingManager {
	return &CoolingManager{
		devices:      make(map[uint64]float32),
		isMonitoring: false,
		MaxTemp:      maxTemp,
	}
}

func (c *CoolingManager) AddDevice(id uint64, temp float32) error {
	_, ok := c.devices[id]
	if !ok {
		return errors.New("CoolingManager AddDevice Error")
	}

	c.devices[id] = temp
	return nil
}

func (c *CoolingManager) monitor() error {
	var errOut error
	go func() {
		err := func() error {
			if !c.isMonitoring {

				return errors.New("CoolingManager monitoring is disabled")
			}
			for {
				time.Sleep(5 * time.Second)
				for k, _ := range c.devices {
					if c.devices[k] > c.MaxTemp {
						return errors.New("Temp over heated")
					}
				}
			}
		}()

		errOut = err
	}()

	return errOut
}

func (c *CoolingManager) Stop() {
	c.isMonitoring = false
}
