package store

//Store provide the storage capacity
type IStore interface {
	//GetMap return a map named group
	GetMap(group []byte)
	//GetQueue return a queue named group
	GetQueue(group []byte)
	//GetList return a stack named group
	GetStack(group []byte)
	//GetArray return a stack named group
	GetArray(group []byte)
	//GetSortedMap return a SortedMap named group
	GetSortedMap(group []byte)
	//GetLevelKeyMap
	GetLevelKeyMap(group []byte)
}
