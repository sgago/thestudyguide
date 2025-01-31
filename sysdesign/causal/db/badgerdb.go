package db

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/dgraph-io/badger/v3"
)

type BadgerDB struct {
	db *badger.DB
}

func NewBadgerDB(path string) (*BadgerDB, error) {
	opts := badger.DefaultOptions(path).
		WithLoggingLevel(badger.ERROR)

	db, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}

	return &BadgerDB{db: db}, nil
}

func (b *BadgerDB) Close() error {
	return b.db.Close()
}

func (b *BadgerDB) Set(key string, value []byte, ttl time.Duration) error {
	return b.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		return txn.SetEntry(e)
	})
}

func (b *BadgerDB) Get(key string) ([]byte, error) {
	var valCopy []byte

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		return err
	})

	return valCopy, err
}

func (b *BadgerDB) Delete(key string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// ToBytes converts a generic value to a byte array.
func ToBytes[T any](value T) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(value); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// FromBytes converts a byte array back to a generic value.
func FromBytes[T any](data []byte) (T, error) {
	var value T

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&value); err != nil {
		return value, err
	}

	return value, nil
}
