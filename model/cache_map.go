package model

import (
	"gin-example/common"
	"sync"
)

type cachedValue struct {
	value interface{}
}
type CacheMap struct {
	mux      *sync.RWMutex
	data     map[string]*cachedValue
	deepCopy func(interface{}) (interface{}, error)
}

func NewCacheMap() *CacheMap {
	return &CacheMap{
		mux:      &sync.RWMutex{},
		data:     make(map[string]*cachedValue),
		deepCopy: common.DeepCopy,
	}
}

func (m *CacheMap) Put(k string, v interface{}) bool {
	defer m.mux.Unlock()
	m.mux.Lock()

	if v == nil {
		return false
	}

	vcpy, err := m.deepCopy(v)
	if err != nil {
		return false
	}

	m.data[k] = &cachedValue{
		value: vcpy,
	}
	return true
}

func (m *CacheMap) Get(k string) interface{} {
	v := m.get(k)
	if v == nil {
		return nil
	}

	if vcpy, err := m.deepCopy(v); err == nil {
		return vcpy
	}

	return nil
}

func (m *CacheMap) get(k string) interface{} {
	defer m.mux.RUnlock()
	m.mux.RLock()
	entry, ok := m.data[k]
	if !ok {
		return nil
	}
	return entry.value
}
