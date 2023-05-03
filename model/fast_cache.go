package model

import (
	"gin-example/util"
	"github.com/VictoriaMetrics/fastcache"
	"sync/atomic"
)

type fastCache struct {
	name    string
	data    *atomic.Pointer[fastcache.Cache]
	encoder Encoder
}

func NewFastCache(name string, encoder Encoder) KVSnapshot {
	atomicPointer := atomic.Pointer[fastcache.Cache]{}
	atomicPointer.Store(nil)

	return &fastCache{
		name:    name,
		data:    &atomicPointer,
		encoder: encoder,
	}
}

type compressed [][]byte

func (s *fastCache) Name() string {
	return s.name
}

func (s *fastCache) Get(k string) (any, error) {
	cache := s.data.Load()
	if cache == nil {
		return nil, nil
	}

	key := util.ToBytes(k)
	var entityByte []byte
	entityByte = cache.Get(entityByte, key)
	if len(entityByte) == 0 {
		return nil, nil
	}

	v := s.encoder.Decode(entityByte)
	return v, nil
}

func (s *fastCache) Set(data []any, f KeyExtractor) (bool, []error) {
	errors := make([]error, 0, 10)
	var length = len(data)
	var dataSize int
	var insertElements int
	var entities = make([]compressed, 0, length)

	for _, entity := range data {
		if entity == nil {
			continue
		}

		var key, err = f(entity)

		if err != nil {
			errors = append(errors, err)
			continue
		}

		var keyByte = util.ToBytes(key)
		entityByte, err := s.encoder.Encode(entity)

		if err != nil {
			errors = append(errors, err)
			continue
		}

		dataSize += len(entityByte)
		var keyValue = compressed{keyByte, entityByte}
		entities = append(entities, keyValue)
	}

	newCache := fastcache.New(dataSize)
	for _, entitySlice := range entities {
		insertElements += 1
		newCache.Set(entitySlice[0], entitySlice[1])
	}

	s.data.Swap(newCache)

	if len(errors) != 0 {
		return false, errors
	}

	return true, nil
}
