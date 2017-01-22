package store

import (
	"errors"
	"github.com/boltdb/bolt"
)

var (
	ErrBucketNotFound  = errors.New("Bucket not found")
	ErrKeyNotFound	   = errors.New("Key not found")
	ErrDoesNotExists   = errors.New("Does Found it")
	ErrFoundIt         = errors.New("Found it")
	ErrExistsInSet     = errors.New("Element already exists in set")
	ErrInvalidId       = errors.New("Element ID can not contain \":\"")
)


type DataBase struct {
	*bolt.DB
}


type (
	//Used for each of the datatypes
	boltBucket struct {
		db *DataBase //the Bolt database
		name []byte //the bucket name
	}

	List 		boltBucket
	Set 		boltBucket
	HashMap 	boltBucket
	KeyValue	boltBucket
	Queue		boltBucket
	Stack		boltBucket
	LinkedMap	boltBucket
)

//Store provide the storage capacity
type IStore interface {
	//GetMap return a map named group
	GetMap(group []byte)
	//GetQueue return a queue named group
	GetQueue(group []byte)
	//GetList return a stack named group
	GetStack(group []byte)
	//GetArray return a stack named group
	GetList(group []byte)
	//GetSet return a set named group
	GetSet(group []byte)
	//GetSortedMap return a SortedMap named group
	GetLinkedMap(group []byte)
	//GetLevelKeyMap
	GetKeyValue(group []byte)
}
