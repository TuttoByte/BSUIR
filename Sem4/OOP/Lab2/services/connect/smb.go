package connect

import "net"

type SMBConnetcion struct {
	session net.Conn
}

func (s *SMBConnetcion) EstablishConnection() error {
	return nil
}
func (s *SMBConnetcion) CloseConnection() error {
	return nil
}
func (s *SMBConnetcion) ShowRemoteFiles() error {
	return nil
}
func (s *SMBConnetcion) DownRemoteloadFile() error {
	return nil
}
func (s *SMBConnetcion) ListRemoteFiles() error {
	return nil
}
