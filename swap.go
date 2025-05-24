package scalingSuffix

import "maps"
import "iter"

func Pivot[Map ~map[K]V, K comparable, V comparable](m Map) map[V]K{
	return maps.Collect(swap[K,V](maps.All(m)))
}

// makes a reverse map
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



