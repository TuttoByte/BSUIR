package conn

import (
	"fmt"
)

type TelegramConnection struct {
	connectionTocken string
}

func NewTelegramConnection(tocken string) *TelegramConnection {
	return &TelegramConnection{}
}

func (t *TelegramConnection) Connect() {
	fmt.Printf("Connected to tocken %s", t.connectionTocken)
}
func (t *TelegramConnection) Disconnect() {
	fmt.Printf("Disconected from tocken %s", t.connectionTocken)

}
func (t *TelegramConnection) Send(msg string) {
	fmt.Printf("Send massage to user by tocken %s")
}
