package managers

import (
	"Lab3/services/monitor"
	"fmt"
)

type Connection interface {
	Connect(logger *monitor.CustomLogger)
	Disconnect(logger *monitor.CustomLogger)
	Send(msg string, logger *monitor.CustomLogger)
}

type ConnectionManager struct {
	conns []Connection
}

func NewConnectionManager(args ...Connection) *ConnectionManager {

	conns := make([]Connection, len(args))

	copy(conns, args)

	return &ConnectionManager{
		conns: conns,
	}

}
func (c *ConnectionManager) SendOverConnection(msg string, logger *monitor.CustomLogger) {
	for _, conn := range c.conns {
		fmt.Println(conn)
		conn.Connect(logger)
		conn.Send(msg, logger)
	}
}
