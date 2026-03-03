package conn

import (
	"Lab3/services/monitor"
	"fmt"
)

type TelegramConnection struct {
	connectionTocken string
}

func NewTelegramConnection(tocken string) *TelegramConnection {
	fmt.Printf(">> Connecting to Telegram by tocken %s...\n", tocken)
	return &TelegramConnection{
		connectionTocken: tocken,
	}
}

func (t *TelegramConnection) Connect(logger *monitor.CustomLogger) {
	logger.Info("Telegram", ">> Connected to tocken %s", t.connectionTocken)
}
func (t *TelegramConnection) Disconnect(logger *monitor.CustomLogger) {
	logger.Info(">> Disconected from tocken %s", t.connectionTocken)

}
func (t *TelegramConnection) Send(msg string, logger *monitor.CustomLogger) {
	logger.Info(">> Send massage to user by tocken %s")
}
