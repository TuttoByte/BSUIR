package control

import (
	"Lab1/interfaces"
	"net"
)

type Client struct {
	id            uint64
	ip            net.IP
	conn          interfaces.Connection
	controlStatus bool
}

func NewClient(id uint64, ip net.IP, conn interfaces.Connection) *Client {
	return &Client{
		id:            id,
		ip:            ip,
		conn:          conn,
		controlStatus: false,
	}
}

func (c *Client) GetID() uint64 {
	return c.id
}
func (c *Client) GetIP() net.IP {
	return c.ip
}
func (c *Client) GetConn() interfaces.Connection {
	return c.conn
}
func (c *Client) GetStatus() bool {
	return c.controlStatus
}

func (c *Client) SetControlStatus(status bool) {
	c.controlStatus = status
}
