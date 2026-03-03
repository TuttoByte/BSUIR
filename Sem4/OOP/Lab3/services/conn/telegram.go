package conn

import (
	"Lab3/services/monitor"
)

type TelegramConnection struct {
	connectionTocken string
}

func NewTelegramConnection(tocken string) *TelegramConnection {

	return &TelegramConnection{
		connectionTocken: tocken,
	}
}

func (t *TelegramConnection) Connect(logger *monitor.CustomLogger) {
	logger.Info("Telegram", ">> Connected to tocken %s", t.connectionTocken)
}
func (t *TelegramConnection) Disconnect(logger *monitor.CustomLogger) {
	logger.Info("Telegram", ">> Disconected from tocken %s", t.connectionTocken)

}
func (t *TelegramConnection) Send(msg string, logger *monitor.CustomLogger) {
	logger.Info("Telegram", ">> Send massage to user by tocken %s")
}
