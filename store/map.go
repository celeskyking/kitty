package store

import (
	"errors"
	"github.com/boltdb/bolt"
)



func NewMap(db *DataBase,id string) (*HashMap,error) {
	name := []byte(id)
	if err := db.Update(func(tx *bolt.Tx)error {
		if _, err := tx.CreateBucketIfNotExists(name);err != nil {
			return errors.New("can not create bucket:"+err.Error())
		}
	});err != nil {
		return nil, err
	}
	return &HashMap{db,name},nil
}

//Put key and value to the disk
func (m *HashMap) Put(key string, value []byte) (err error){
	if m.name == nil {
		return ErrDoesNotExists
	}
	return m.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.Put([]byte(key),value)
	})
}

//Get the value of key
func (m *HashMap) Get(key string) (result []byte,err error) {
	if m.name == nil {
		return nil,ErrDoesNotExists
	}
	return result,m.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		byteVal := bucket.Get([]byte(key))
		if byteVal == nil {
			return ErrKeyNotFound
		}
		result = byteVal
		return nil
	})
}

//Has
func (m *HashMap) Has(key string) (found bool,err error) {
	if m.name == nil {
		return false,ErrDoesNotExists
	}
	return found, m.db.View(func(tx *bolt.Tx) error{
		bucket := tx.Bucket(m.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		byteVal := bucket.Get([]byte(key))
		if byteVal != nil {
			found = true
		}
		return nil
	})
}

func(m *HashMap) Values()(results [][]byte,err error) {
	if m.name == nil {
		return nil, ErrDoesNotExists
	}
	return results, m.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_,result []byte) error {
			results = append(results,result)
			return nil
		})
	})
}


func (m *HashMap) Keys() (keys []string, err error){
	if m.name == nil {
		return nil, ErrDoesNotExists
	}
	return keys, m.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(key,_ []byte) error {
			keys = append(keys,string(key))
			return nil
		})
	})
}



