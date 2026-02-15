package cmdManager

import "sync"

type ServerManager struct {
	wg      sync.WaitGroup
	configs map[string]string
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
