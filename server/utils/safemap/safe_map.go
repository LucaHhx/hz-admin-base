package safemap

import "sync"

// SafeMap 并发安全的泛型映射
type SafeMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V

	initialized bool
}

// NewSafeMap 创建新的并发安全Map
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		m: make(map[K]V),
	}
}

func (s *SafeMap[K, V]) MSet(data map[K]V) {
	s.Lock()
	defer s.Unlock()
	for key, value := range data {
		s.m[key] = value
	}
}

func (s *SafeMap[K, V]) Store(key K, value V) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *SafeMap[K, V]) Load(key K) (value V, ok bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok = s.m[key]
	return
}

// LoadOrStore 存在时加载，不存在时存储
func (s *SafeMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	s.Lock()
	defer s.Unlock()

	actual, loaded = s.m[key]
	if !loaded {
		s.m[key] = value
		actual = value
	}
	return
}

func (s *SafeMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	s.Lock()
	defer s.Unlock()

	value, loaded = s.m[key]
	if loaded {
		delete(s.m, key)
	}
	return
}

func (s *SafeMap[K, V]) Delete(keys ...K) {
	s.Lock()
	defer s.Unlock()
	for _, key := range keys {
		delete(s.m, key)
	}
}

func (s *SafeMap[K, V]) Pop(key K) (v V, exists bool) {
	s.Lock()
	defer s.Unlock()
	v, exists = s.m[key]
	delete(s.m, key)
	return
}

func (s *SafeMap[K, V]) Swap(key K, value V) (actual V, loaded bool) {
	s.Lock()
	defer s.Unlock()

	actual, loaded = s.m[key]
	s.m[key] = value
	return
}

func (s *SafeMap[K, V]) MSwap(m map[K]V) {
	s.Lock()
	defer s.Unlock()

	s.m = make(map[K]V)
	for k, v := range m {
		s.m[k] = v
	}
	return
}

func (s *SafeMap[K, V]) Items() map[K]V {
	s.RLock()
	defer s.RUnlock()

	tmp := make(map[K]V)
	for k, v := range s.m {
		tmp[k] = v
	}
	return tmp
}

func (s *SafeMap[K, V]) Keys() []K {
	s.RLock()
	defer s.RUnlock()

	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *SafeMap[K, V]) Values() []V {
	s.RLock()
	defer s.RUnlock()

	values := make([]V, 0, len(s.m))
	for _, v := range s.m {
		values = append(values, v)
	}
	return values
}

// Clear 清空映射
func (s *SafeMap[K, V]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[K]V)
	s.initialized = false
}

// Size 获取元素数量
func (s *SafeMap[K, V]) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.m)
}

func (s *SafeMap[K, V]) Init(list []V, getKey func(v V) (K, bool)) {
	s.Lock()
	defer s.Unlock()
	for _, v := range list {
		if k, ok := getKey(v); ok {
			s.m[k] = v
		}
	}
	s.initialized = true
}

// IsInitialized 是否已经初始化
func (s *SafeMap[K, V]) IsInitialized() bool {
	s.RLock()
	defer s.RUnlock()
	return s.initialized
}
