package transaction

import (
	"errors"
	"sync"
	"time"
)

var transactions = []Transaction{
	{Id: 1, Title: "Title 1", Amount: 650.75, Type: 2, CreatedAt: time.Now()},
	{Id: 2, Title: "Title 2", Amount: 50.50, Type: 1, CreatedAt: time.Now()},
	{Id: 3, Title: "Title 3", Amount: 73.75, Type: 1, CreatedAt: time.Now()},
}

type memoryDB struct {
	sync.RWMutex
	Data []Transaction
}

var db *memoryDB

func GetDB() *memoryDB {
	if db == nil {
		db = &memoryDB{
			Data: transactions,
		}
	}

	return db
}

var (
	ErrIdNotFound = errors.New("id not found")
)

func (mdb *memoryDB) FindAll() ([]Transaction, error) {
	mdb.RLock()
	defer mdb.RUnlock()

	return mdb.Data, nil
}

func (mdb *memoryDB) FindById(id int64) (*Transaction, error) {
	mdb.RLock()
	defer mdb.RUnlock()

	for _, transaction := range mdb.Data {
		if transaction.Id == id {
			return &transaction, nil
		}
	}

	return nil, ErrIdNotFound
}

func (mdb *memoryDB) Create(newTransaction Transaction) (Transaction, error) {
	mdb.Lock()
	defer mdb.Unlock()

	newId := len(mdb.Data) + 1

	newTransaction.Id = int64(newId)
	newTransaction.CreatedAt = time.Now()

	mdb.Data = append(mdb.Data, newTransaction)

	return newTransaction, nil
}
