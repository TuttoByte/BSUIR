package conn

import (
	"Lab3/services/monitor"
	"fmt"
)

type SMTPConnection struct {
	server string
}

func NewSMTPConnection(ip string) *SMTPConnection {
	fmt.Printf(">> Connecting to SMTP server %s...\n", ip)
	return &SMTPConnection{
		server: ip,
	}
}

func (s *SMTPConnection) Connect(logger *monitor.CustomLogger) {
	logger.Info(">> Connecting to SMTP server %s...\n", s.server)
}

func (s *SMTPConnection) Disconnect(logger *monitor.CustomLogger) {
	logger.Info(">> Disconnect from SMTP server %s...\n", s.server)
}

func (s *SMTPConnection) Send(msg string, logger *monitor.CustomLogger) {
	logger.Info(">> SMTP sended with connection over server " + s.server)
}
