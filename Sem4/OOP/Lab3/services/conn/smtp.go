package conn

import (
	"Lab3/services/monitor"
)

type SMTPConnection struct {
	server string
}

func NewSMTPConnection(ip string) *SMTPConnection {
	return &SMTPConnection{
		server: ip,
	}
}

func (s *SMTPConnection) Connect(logger *monitor.CustomLogger) {
	logger.Info("SMTP", ">> Connecting to server %s...\n", s.server)
}

func (s *SMTPConnection) Disconnect(logger *monitor.CustomLogger) {
	logger.Info("SMTP", ">> Disconnect from SMTP server %s...\n", s.server)
}

func (s *SMTPConnection) Send(msg string, logger *monitor.CustomLogger) {
	logger.Info("SMTP", ">>sended with connection over server "+s.server)
}
