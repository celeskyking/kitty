package store

import (
	"errors"

	"github.com/boltdb/bolt"
)


//Map the data struct to map
type Map struct {
	conn  *bolt.DB
	group string
}

//Put key and value to the disk
func (m *Map) Put(key string, value []byte) {
	m.conn.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(m.group)).Put([]byte(key), value)
		return err
	})
}

//Get the value of key
func (m *Map) Get(key string) []byte {
	var value []byte
	m.conn.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte(m.group)).Get([]byte(key))
		return nil
	})
	return value
}

//Keys return the keys of the limit
func (m *Map) Keys(size int) [][]byte {
	i := 0
	keys := make([][]byte, 0)
	m.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(m.group))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			keys = append(keys, k)
			i++
		}
		return nil
	})
	return keys
}

//Entry is a key-value pair of map
type Entry struct {
	Key   string
	Value []byte
}
