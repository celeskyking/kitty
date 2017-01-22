package store

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/celeskyking/kitty"
)

//NewSet return a set
func NewSet(db *DataBase, id string)(*Set,error) {
	name := []byte(id)
	if err := db.Update(func(tx *bolt.Tx) error{
		if _, err := tx.CreateBucketIfNotExists(name);err != nil {
			return errors.New("can not create bucket:"+err.Error())
		}
	});err != nil {
		return nil, err
	}
	return &Set{db,name},nil
}

func (s *Set) Add(value []byte) error {
	if s.name == nil {
		return ErrDoesNotExists
	}
	exists,err := s.Has(value)
	if err != nil {
		return err
	}
	if exists {
		return ErrExistsInSet
	}
	return s.db.Update(func(tx *bolt.Tx)error{
		bucket := tx.Bucket(s.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		n, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		return bucket.Put(kitty.Uint64ToBytes(n),value)
	})
}

func (s *Set) Has(value []byte) (exists bool,err error){
	if s.name  == nil {
		return false,ErrDoesNotExists
	}
	return exists, s.db.Update(func(tx *bolt.Tx)error {
		bucket := tx.Bucket(s.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_,v []byte)error{
			if kitty.Md5Hex(v) == kitty.Md5Hex(value){
				exists = true
				return ErrFoundIt
			}
			return nil
		})
		return nil
	})
}

func (s *Set) GetAll() (values [][]byte,err error){
	if s.name == nil {
		return nil, ErrDoesNotExists
	}
	return values,s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_,value []byte) error {
			values = append(values,value)
			return nil
		})
	})
}



func (s *Set) Del(value []byte) error {
	if s.name == nil {
		return ErrDoesNotExists
	}
	return s.db.Update(func(tx *bolt.Tx) error{
		bucket := tx.Bucket(s.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		var foundKey []byte
		bucket.ForEach(func(byteKey,byteValue []byte) error {
			if kitty.Md5Hex(byteValue) == kitty.Md5Hex(value){
				foundKey = byteValue
				return ErrFoundIt
			}
			return nil
		})
		return bucket.Delete(foundKey)
	})
}


func (s *Set) Remove() error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(s.name)
	})
	s.name = nil
	return err
}

func(s *Set) Clear() error {
	if s.name == nil {
		return ErrDoesNotExists
	}
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(key,_ []byte) error {
			return bucket.Delete(key)
		})
	})
}