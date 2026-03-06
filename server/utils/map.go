package utils

func MapFilter[T comparable, V any](data map[T]V, filter func(T, V) bool) map[T]V {
	res := make(map[T]V)
	for k, v := range data {
		if filter(k, v) {
			res[k] = v
		}
	}
	return res
}

func MapFilterKey[T comparable, V any](data map[T]V, filter func(T, V) bool) []T {
	var res []T
	for k, v := range data {
		if filter(k, v) {
			res = append(res, k)
		}
	}
	return res
}
