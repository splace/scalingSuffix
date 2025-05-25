package scalingSuffix

import "golang.org/x/exp/constraints"
import "fmt"
import "strconv"
import "strings"
import "slices"

// SIUnits's are SI's with metric units suffixes added.
type SIUnits[N constraints.Float | constraints.Integer] struct{SI[N];fmt.Stringer}

func NewSIUnits[N constraints.Float | constraints.Integer](v N, f fmt.Stringer) SIUnits[N]{
	return SIUnits[N]{SI[N]{v},f}
}

//type Dimension uint8

//const (
//	Length Dimension iota
//	Masss
//	Time
//)

func DimString(d int8, s string)(_ string){
	switch d{
	case 0:
		return
	case 1:
		return s
	default:
		return s+strconv.Itoa(int(d))
	}
}

type Dims struct{M,L,T int8}

func (d Dims) String()string{
	return strings.Join(
		slices.DeleteFunc(
			[]string{
				DimString(d.M,"kg"),
				DimString(d.L,"m"),
				DimString(d.T,"s"),
			},
			func(s string)bool{return s==""},
		),
		".",
	)
}


//type BaseUnit uint8

//const (
//	Temperature BaseUnit iota	
//	Voltage	
//	Current
//)

//type CompositeUnit uint8

//const(
//	Energy CompositeUnit
//	Velocity
//	Acceleration
//	Force
//	Power
//	Pressure
//)



