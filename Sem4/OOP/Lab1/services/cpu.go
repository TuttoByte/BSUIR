package services

import (
	"errors"
	"github.com/northwindlight/cputemp"
)

type CPU struct {
	cores       uint64
	threads     uint64
	nm          uint64
	isStarted   bool
	temperature float32
	boost       bool
	Model       string
}

func NewCPU(cores, threads, nm uint64, model string) *CPU {
	return &CPU{
		cores:   cores,
		threads: threads,
		nm:      nm,
		Model:   model,
	}
}

func (c *CPU) getCPUSystemTemp() float64 {
	temp, err := cputemp.GetCPUTemperature()
	if err != nil {
		return 0.0
	}
	return temp
}

func (c *CPU) Start() error {
	if c.isStarted {
		return errors.New("cpu already started")
	}
	c.isStarted = true
	return nil
}

func (c *CPU) Stop() error {
	if !c.isStarted {
		return errors.New("cpu already stopped")
	}
	c.isStarted = false
	return nil
}
