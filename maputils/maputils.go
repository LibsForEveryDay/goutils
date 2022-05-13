package maputils

import (
	"fmt"
	"sort"
)

type GenericPair[K comparable, V any] struct {
	Key   K
	Value V
}

type GenericPairList[K comparable, V any] []GenericPair[K, V]

// Sort sorts the overall map[K]V by key or by value. It returns GenericPairList[K, V].
func Sort[M map[K]V, K comparable, V any](
	m M,
	fKey func(k1, k2 K) bool,
	fVal func(v1, v2 V) bool,
	reverse bool,
) GenericPairList[K, V] {
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

// Sort sorts the overall map[K]V by key. It returns GenericPairList[K, V].
func SortByKey[M map[K]V, K comparable, V any](m M, fKey func(k1, k2 K) bool, reverse bool) GenericPairList[K, V] {
	return Sort(m, fKey, nil, reverse)
}

// Sort sorts the overall map[K]V by value. It returns GenericPairList[K, V].
func SortByVal[M map[K]V, K comparable, V any](m M, fVal func(v1, v2 V) bool, reverse bool) GenericPairList[K, V] {
	return Sort(m, nil, fVal, reverse)
}

func PrintGenericList[K comparable, V any](pairs GenericPairList[K, V]) {
	fmt.Printf("%T:\n", pairs)
	for i, pair := range pairs {
		fmt.Printf("%2s- Index: %d, Key: %+v, Value: %+v\n", "", i, pair.Key, pair.Value)
	}
}
