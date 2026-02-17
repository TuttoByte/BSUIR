package connect

import "net"

type FTPConnetcion struct {
	session net.Conn
}

func (f *FTPConnetcion) EstablishConnection() error {
	return nil
}
func (f *FTPConnetcion) CloseConnection() error {
	return nil
}
func (f *FTPConnetcion) ShowRemoteFiles() error {
	return nil
}
func (f *FTPConnetcion) DownRemoteloadFile() error {
	return nil
}
func (f *FTPConnetcion) ListRemoteFiles() error {
	return nil
}
