package contracts

type DataBase interface {
	AddUser(name, hash string) error
	GetUser(name string) (string, error)
	Close() error
}
