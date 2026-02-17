package managres

import (
	"Lab1/interfaces"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"time"
)

type BackupManager struct {
	storageData *bytes.Buffer
	period      uint64
	isPeriodic  bool
}

type BackupColection[T any] struct {
	data       []T
	presethash []string
	log.Logger
}

func (b *BackupColection[T]) CheckHash() error {
	var err error
	go func() {
		for {
			var i int
			time.Sleep(time.Duration(5) * time.Second)
			buffer := bytes.NewBuffer(nil)
			err := binary.Write(buffer, binary.LittleEndian, b.data)
			if err != nil {
				err = errors.New(fmt.Sprint("binary write error: ", err))
				return
			}
			hasher := sha1.New()
			hasher.Write(buffer.Bytes())
			sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
			if sha == b.presethash[i] {
				return
			}

			if i == len(b.data) {
				i = 0
			}

		}
	}()
	return err
}

func NewBackupManager(period uint64, isPeriod bool) *BackupManager {
	return &BackupManager{
		storageData: bytes.NewBuffer(make([]byte, 0)),
		period:      period,
		isPeriodic:  isPeriod,
	}
}

func (b *BackupManager) Save(disks []interfaces.Device) error {
	for _, disk := range disks {
		err := binary.Write(b.storageData, binary.LittleEndian, disk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BackupManager) proceesPeriod(disks []interfaces.Device) error {
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

func (b *BackupManager) StartPeriod(disk []interfaces.Device) error {
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

func (b *BackupManager) GetStatus() bool {
	return b.isPeriodic
}
