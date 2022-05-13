package maputils

import (
	"sort"
)

type GenericPair[K comparable, V any] struct {
	Key   K
	Value V
}

// Sort sorts the overall map[K]V by key or by value. It returns GenericPair[K]V.
func Sort[M map[K]V, K comparable, V any](
	m M,
	fKey func(k1, k2 K) bool,
	fVal func(v1, v2 V) bool,
	reverse bool,
) []GenericPair[K, V] {
	var res []GenericPair[K, V] = make([]GenericPair[K, V], 0, len(m))

	for key, val := range m {
		res = append(res, GenericPair[K, V]{key, val})
	}

	var status bool
	sort.Slice(res, func(i, j int) bool {
		if fVal != nil {
			status = fVal(res[i].Value, res[j].Value)
		} else {
			status = fKey(res[i].Key, res[j].Key)
		}
		if reverse {
			status = !status
		}
		return status
	})

	return res
}

// Sort sorts the overall map[K]V by key. It returns GenericPair[K]V.
func SortByKey[M map[K]V, K comparable, V any](m M, fKey func(k1, k2 K) bool, reverse bool) []GenericPair[K, V] {
	return Sort(m, fKey, nil, reverse)
}

// Sort sorts the overall map[K]V by value. It returns GenericPair[K]V.
func SortByVal[M map[K]V, K comparable, V any](m M, fVal func(v1, v2 V) bool, reverse bool) []GenericPair[K, V] {
	return Sort(m, nil, fVal, reverse)
}
