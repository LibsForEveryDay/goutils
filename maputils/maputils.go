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
//
//  Example:
//  	func main() {
//  		data := map[string]int{
//  			"a": 3, "b": 1, "c": 5,
//  			"d": 0, "e": 2, "f": 4,
//  		}
//
//  		fmt.Println("Original:", data)
//
//  		res := maputils.Sort(data, func(k1, k2 maputils.GenericPair[string, int]) bool {
//  			return k1.Key < k2.Key
//  		}, false)
//
//  		fmt.Println("\nSorted by key:")
//  		maputils.PrintGenericList(res)
//
//  		res = maputils.Sort(data, func(k1, k2 maputils.GenericPair[string, int]) bool {
//  			return k1.Value < k2.Value
//  		}, false)
//
//  		fmt.Println("\nSorted by value:")
//  		maputils.PrintGenericList(res)
//  	}
//
func Sort[M map[K]V, K comparable, V any](
	m M,
	f func(p1, p2 GenericPair[K, V]) bool,
	reverse bool,
) GenericPairList[K, V] {
	var res []GenericPair[K, V] = make([]GenericPair[K, V], 0, len(m))

	for key, val := range m {
		res = append(res, GenericPair[K, V]{key, val})
	}

	var status bool
	sort.Slice(res, func(i, j int) bool {
		status = f(res[i], res[j])
		if reverse {
			status = !status
		}
		return status
	})

	return res
}

func PrintGenericList[K comparable, V any](pairs GenericPairList[K, V]) {
	fmt.Printf("%T:\n", pairs)
	for i, pair := range pairs {
		fmt.Printf("%2s- Index: %d, Key: %+v, Value: %+v\n", "", i, pair.Key, pair.Value)
	}
}
