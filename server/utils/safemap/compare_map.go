package safemap

type CompareMap[K comparable, V comparable] struct {
	*SafeMap[K, V]
}

func NewCompareMap[K comparable, V comparable]() *CompareMap[K, V] {
	return &CompareMap[K, V]{
		SafeMap: NewSafeMap[K, V](),
	}
}

// CompareAndSwap 比较并交换
func (m *CompareMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	m.Lock()
	defer m.Unlock()
	if m.m[key] == old {
		m.m[key] = new
		swapped = true
	}
	return
}

// CompareAndDelete 比较并删除
func (m *CompareMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	m.Lock()
	defer m.Unlock()
	if m.m[key] == old {
		delete(m.m, key)
		deleted = true
	}
	return
}
