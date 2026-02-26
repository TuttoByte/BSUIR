package managers

type Connection interface {
	Connect()
	Disconnect()
	Send(msg string)
}

type ConnectionManager struct {
	conns []Connection
}

func NewConnectionManager(args ...Connection) *ConnectionManager {

	conns := make([]Connection, len(args))

	for _, conn := range args {
		conns = append(conns, conn)
	}

	return &ConnectionManager{
		conns: conns,
	}

}
func (c *ConnectionManager) SendOverConnection(msg string) {
	for _, conn := range c.conns {
		conn.Connect()
		conn.Send(msg)
	}
}
