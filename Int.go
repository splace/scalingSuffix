package scalingSuffix

import "golang.org/x/exp/constraints"
import "fmt"

type integer[N constraints.Integer] struct{
	Value N
}

// Int is a fmt.Stringer writing numbers comma separated in blocks of 1000's
type Int[N constraints.Signed] integer[N]

func NewInt[N constraints.Signed](v N) Int[N]{
	return Int[N]{v}
}

func (i Int[N]) String()string{
	if i.Value>=0 {
		return Uint[uint64]{uint64(i.Value)}.String()
	}
	return "-"+Uint[uint64]{uint64(-i.Value)}.String()
}

// Uint is a fmt.Stringer writing numbers comma separated in blocks of 1000's
type Uint[N constraints.Unsigned] integer[N]

func NewUint[N constraints.Unsigned](v N) Uint[N]{
	return Uint[N]{v}
}

func (i Uint[N]) String()string{
	return SpaceSep3(fmt.Sprint(i.Value))
}

