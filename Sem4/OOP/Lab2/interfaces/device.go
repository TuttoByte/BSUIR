package interfaces

import (
	"net"
)

type Device interface {
	GetType() string
	GetUAPI() net.Listener
	GetMaxTemp() float32
	Connect() error
	Disconnect() error
}
