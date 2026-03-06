package safemap

import (
	cmap "github.com/orcaman/concurrent-map/v2"
)

type CMap[K comparable, V any] struct {
	*cmap.ConcurrentMap[K, V]
}

func NewMapCache[V any]() *CMap[string, V] {
	data := cmap.New[V]()
	cache := &CMap[string, V]{
		ConcurrentMap: &data,
	}
	return cache
}

func NewMapCacheUint[V any]() *CMap[uint, V] {
	data := cmap.NewWithCustomShardingFunction[uint, V](func(key uint) uint32 {
		return uint32(key)
	})
	return &CMap[uint, V]{
		ConcurrentMap: &data,
	}
}

func (mp *CMap[K, V]) Keys() []K {
	var keys []K
	for key := range mp.Items() {
		keys = append(keys, key)
	}
	return keys
}

func (mp *CMap[K, V]) Values() []V {
	var values []V
	for _, value := range mp.Items() {
		values = append(values, value)
	}
	return values
}

func (mp *CMap[K, V]) Find(f func(K, V) bool) (V, bool) {
	for key, value := range mp.Items() {
		if f(key, value) {
			return value, true
		}
	}
	return *new(V), false
}

func (mp *CMap[K, V]) Init(list []V, getKey func(v V) (K, bool)) {
	for _, v := range list {
		if k, ok := getKey(v); ok {
			mp.Set(k, v)
		}
	}
}
