package raft

import (
	"github.com/celeskyking/kitty/store"
	"github.com/hashicorp/raft"
	"sync"
)

type StorageFSM struct {
	mux 	*sync.Mutex
	rs	store.IStore
}



func (s *StorageFSM) Apply(log *raft.Log) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
}