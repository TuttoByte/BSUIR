package interfaces

type Connection interface {
	EstablishConnection() error
	CloseConnection() error
	ShowRemoteFiles() error
	ListRemoteFiles() error
	DownRemoteloadFile() error
}
