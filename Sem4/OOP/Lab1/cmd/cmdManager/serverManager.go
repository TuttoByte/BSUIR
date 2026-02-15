package cmdManager

type ServerManager struct {
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
	go task()
}
