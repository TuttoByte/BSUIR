package cmdManager

import (
	"Lab1/interfaces"
	"Lab1/services"
	"Lab1/services/managres"
	"Lab1/services/network"
	"Lab1/services/network/control"
	"sync"
)

type ServerManager struct {
	wg      sync.WaitGroup
	configs map[string]string
	devices []interfaces.Device
	users   []control.Client

	//Manager Sections
	NetManager     *network.NetworkManager
	AcessControl   *services.AccessControl
	CollingManager *managres.CoolingManager
	BackupManager  *managres.BackupManager
	AlertManager   *services.AlertManager
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		configs:       make(map[string]string),
		devices:       make([]interfaces.Device, 0),
		users:         make([]control.Client, 0),
		AcessControl:  services.NewAccessControl(),
		BackupManager: managres.NewBackupManager(3, true),
		AlertManager:  services.NewAlertManager(),
		//CollingManager: managres.NewCoolingManager(),
	}
}

func (sm *ServerManager) GetConfig(key string) string {
	return sm.configs[key]
}
func (sm *ServerManager) SetConfig(key string, value string) {
	sm.configs[key] = value
}

func (sm *ServerManager) AddTask(task func()) {
	sm.wg.Add(1)
	go func() {
		defer sm.wg.Done()
		task()
	}()
}

func (sm *ServerManager) Wait() {
	sm.wg.Wait()
}

func (sm *ServerManager) GetDevices() []interfaces.Device {
	return sm.devices
}
