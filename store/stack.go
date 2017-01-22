package store

import (
	"errors"
	"github.com/boltdb/bolt"
)

func NewStack(db *DataBase, id string)(*Set,error){
	name := []byte(id)
	if err := db.Update(func(tx *bolt.Tx) error {
		if _,err := tx.CreateBucketIfNotExists(name);err != nil {
			return errors.New("can not create bucket:"+err.Error())
		}
	});err != nil {
		return nil, err
	}
	return &Set{db,name},nil
}

