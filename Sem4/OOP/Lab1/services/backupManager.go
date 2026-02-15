package services

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"
)

type BackupManager struct {
	storageData *bytes.Buffer
	period      uint64
	isPeriodic  bool
}

func NewBackupManager(period uint64, isPeriod bool) *BackupManager {
	return &BackupManager{
		storageData: bytes.NewBuffer(make([]byte, 0)),
		period:      period,
		isPeriodic:  isPeriod,
	}
}

func (b *BackupManager) Save(disks []Disk) error {
	for _, disk := range disks {
		err := binary.Write(b.storageData, binary.LittleEndian, disk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BackupManager) proceesPeriod(disks []Disk) error {
	var err error
	go func() {
		time.Sleep(time.Duration(b.period) * time.Second)
		if !b.isPeriodic {
			err = errors.New("periodic is disabled")
			return
		}
		err = func() error {
			err := b.Save(disks)
			if err != nil {
				return err
			}
			return nil
		}()
	}()
	return err
}

func (b *BackupManager) StartPeriod(disk []Disk) error {
	if b.isPeriodic {
		return errors.New("periodic save is already enabled")
	}
	b.isPeriodic = true
	err := b.proceesPeriod(disk)
	if err != nil {
		return err
	}
	return nil
}

func (b *BackupManager) StopPeriod() error {
	b.isPeriodic = false
	return nil
}
