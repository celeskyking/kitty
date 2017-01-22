package store

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/celeskyking/kitty"
)

//NewList return a list
func NewList(db *DataBase,id string) (*List,error) {
	name := []byte(id)
	if err := db.Update(func(tx *bolt.Tx)error{
		if _, err := tx.CreateBucketIfNotExists(name);err != nil {
			return errors.New("Could not create bucket:"+err.Error())
		}
		return nil
	});err != nil {
		return nil, err
	}
	return &List{db,name},nil
}

//Add value to the list
func (l *List) Add(value []byte) error {
	if l.name == nil {
		return ErrDoesNotExists
	}
	return l.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(l.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		n, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		return bucket.Put(kitty.Uint64ToBytes(n),[]byte(value))
	})
}

func (l *List) GetAll() (results [][]byte,err error) {
	if l.name == nil {
		return nil, ErrDoesNotExists
	}
	return results,l.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(l.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_,value[]byte)error{
			results = append(results,value)
			return nil
		})
	})
}

//GetLast return the last value of the list
func (l *List) GetLast() (result []byte,err error){
	if l.name == nil {
		return nil,ErrDoesNotExists
	}
	return result,l.db.View(func(tx *bolt.Tx) error{
		bucket := tx.Bucket(l.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		cursor := bucket.Cursor()
		_, result = cursor.Last()
		return nil
	})
}


func(l *List) GetLastN(n int) (results [][]byte,err error){
	if l.name == nil {
		return nil,ErrDoesNotExists
	}
	return results,l.db.View(func(tx *bolt.Tx) error{
		var array [][]byte
		bucket := tx.Bucket(l.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		cursor := bucket.Cursor()
		index := 0
		for k,v := cursor.Last();k!=nil && index < n;k,v = cursor.Prev(){
			array = append(array,v)
			index++
		}
		if len(array) < n {
			return errors.New("too few items in list")
		}
		return nil
	})
}

func(l *List) Remove() error {
	err := l.db.Update(func(tx *bolt.Tx)error{
		return tx.DeleteBucket(l.name)
	})
	l.name = nil
	return err
}


func (l *List) Clear() error {
	if l.name == nil {
		return ErrDoesNotExists
	}
	return l.db.Update(func(tx *bolt.Tx)error{
		bucket := tx.Bucket(l.name)
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(key,_ []byte) error{
			return bucket.Delete(key)
		})
	})
}


