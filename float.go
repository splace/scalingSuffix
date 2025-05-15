package scalingSuffix

import "golang.org/x/exp/constraints"
import "strconv"
import "strings"


type float[N constraints.Float] struct{
	Value N
}

type Float[N constraints.Float] float[N]

func NewFloat[N constraints.Float](v N) Float[N]{
	return Float[N]{v}
}

// Float is a fmt.Stringer writing numbers separated in blocks of 1000
func (f Float[N]) String()string{
	if f.Value<0 {
		ss:=strings.SplitN(strconv.FormatFloat(float64(-f.Value), 'f', -1, 64),".",2)
		ss[0]=SepEvery3(ss[0])
		return "-"+strings.Join(ss,".")
	}
	ss:=strings.SplitN(strconv.FormatFloat(float64(f.Value), 'f', -1, 64),".",2)
	ss[0]=SepEvery3(ss[0])
	return strings.Join(ss,".")
}

