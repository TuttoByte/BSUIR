package cmdManager

import (
	"Lab1/interfaces"
	"Lab1/services"
	"Lab1/services/managres"
	"Lab1/services/network"
	"sync"
)

type ServerManager struct {
	wg             sync.WaitGroup
	configs        map[string]string
	devices        []interfaces.Device
	netManager     *network.NetworkManager
	acessControl   *services.AccessControl
	collingManafer *managres.CoolingManager
	backupManager  *managres.BackupManager
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		configs: make(map[string]string),
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
