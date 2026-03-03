package managers

import (
	"Lab3/services/classes"
	infastruct "Lab3/services/infrastruct"
	"Lab3/services/monitor"
	"crypto/sha256"
	"gopkg.in/yaml.v2"
	"hash"
)

type DbManager struct {
	db        *infastruct.RandomSQLDatabase
	hash      hash.Hash
	hashTable map[string]byte
}

func NewDbManager(db *infastruct.RandomSQLDatabase) *DbManager {
	return &DbManager{
		db:        db,
		hash:      sha256.New(),
		hashTable: make(map[string]byte),
	}
}
func (d *DbManager) computateHash(req string) (bool, error) {
	_, err := d.hash.Write([]byte(req))
	if err != nil {
		return false, err
	}
	hsum := d.hash.Sum(nil)
	_, ok := d.hashTable[string(hsum)]
	return ok, nil
}

func (d *DbManager) AddOrderToDataBase(order classes.Order, total float64, logger *monitor.CustomLogger) (bool, error) {
	buffer, err := yaml.Marshal(order)
	if err != nil {
		return false, err
	}
	ok, err := d.computateHash(string(buffer))
	if err != nil {
		return false, err
	}
	if ok {
		logger.Info("Already in database")
		return true, nil
	}
	err = d.db.SaveOrder(order, total)
	if err != nil {
		return false, err
	}
	return false, nil
}
