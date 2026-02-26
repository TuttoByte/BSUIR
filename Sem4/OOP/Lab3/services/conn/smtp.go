package conn

import "fmt"

type SMTPConnection struct {
	server string
}

func NewSMTPConnection(ip string) *SMTPConnection {
	fmt.Printf(">> Connecting to SMTP server %s...\n", ip)
	return &SMTPConnection{
		server: ip,
	}
}

func (s *SMTPConnection) Connect() {
	fmt.Printf(">> Connecting to SMTP server %s...\n", s.server)
}

func (s *SMTPConnection) Disconnect() {
	fmt.Printf(">> Disconnect from SMTP server %s...\n", s.server)
}

func (s *SMTPConnection) Send(msg string) {
	fmt.Println("Sending over smtp")
	fmt.Println(msg)
}
