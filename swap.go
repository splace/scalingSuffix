package scalingSuffix

import "maps"
import "iter"

// Pivot returns a new map, made from the provided maps key/values swapped (values thus need to be comparable) 
func Pivot[Map ~map[K]V, K, V comparable](m Map) map[V]K{
	return maps.Collect(swap[K,V](maps.All(m)))
}

// swaps a Seq2's pair 
func swap[V1,V2 any](in iter.Seq2[V1,V2])iter.Seq2[V2,V1]{
	return func(yield func(V2, V1) bool) {
 		next, stop := iter.Pull2(in)
   		defer stop()
		for {
			v1, v2, ok1 := next()
			if !ok1 {
				return
			}
			if !yield(v2, v1) {
				return
	   		}
	 	}  
	}
}



