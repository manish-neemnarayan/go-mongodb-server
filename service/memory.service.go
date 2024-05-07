package service

import (
	"fmt"

	"github.com/manish-neemnarayan/go-mongodb-server/types"
)

type MemoryDBer interface {
	Post(*types.MemoryData) (*types.MemoryData, error)
	Get(string) (*types.MemoryData, error)
}

type MemoryDB struct {
	store map[string]string
}

func NewMemoryDB() *MemoryDB {
	fmt.Println("In-memory DB is started")
	return &MemoryDB{
		store: make(map[string]string),
	}
}

func (m *MemoryDB) Post(data *types.MemoryData) (*types.MemoryData, error) {
	m.store[data.Key] = data.Value
	return &types.MemoryData{
		Key:   data.Key,
		Value: m.store[data.Key],
	}, nil
}

func (m *MemoryDB) Get(key string) (*types.MemoryData, error) {
	data, ok := m.store[key]
	if !ok {
		return &types.MemoryData{}, fmt.Errorf("key not found")
	}

	return &types.MemoryData{
		Key:   key,
		Value: data,
	}, nil
}
