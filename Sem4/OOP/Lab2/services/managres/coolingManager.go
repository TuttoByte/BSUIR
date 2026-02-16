package managres

import (
	"Lab1/interfaces"
	"errors"
	"math/rand"
	"net"
	"time"
)

type CoolingDevice struct {
	uapi    net.Listener
	MaxTemp float32
}

type CoolingManager struct {
	devices      map[string]net.Listener
	mSpace       *[]interfaces.Device
	MaxTempList  map[string]float32
	isMonitoring bool
}

func NewCoolingManager(maxTempList map[string]float32) *CoolingManager {
	return &CoolingManager{
		devices:      make(map[string]net.Listener),
		isMonitoring: false,
		MaxTempList:  maxTempList,
	}
}

func (c *CoolingManager) AddDevice(dev interfaces.Device) error {
	_, ok := c.devices[dev.GetType()]
	if !ok {
		return errors.New("CoolingManager AddDevice Error")
	}

	c.devices[dev.GetType()] = dev.GetUAPI()
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
				for _, dev := range *c.mSpace {
					if c.getTemperature(dev.GetUAPI()) > c.MaxTempList[dev.GetType()] {
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

func (c *CoolingManager) Start() {
	c.isMonitoring = true
	if len(*c.mSpace) == 0 {
		return
	}
	go func() {
		_ = c.monitor()
	}()
}

func (c *CoolingManager) AddMonitorSpace(devices []interfaces.Device) error {
	c.mSpace = &devices
	return nil
}
func (c *CoolingManager) getTemperature(uapi net.Listener) float32 {
	return rand.Float32() * 500
}
func (c *CoolingManager) DeviceKill(dev interfaces.Device) error {
	err := dev.Disconnect()
	if err != nil {
		err = dev.GetUAPI().Close()
		if err != nil {
			return err
		}
	}
	return nil
}
