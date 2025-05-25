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

var (
	Energy = Derived{1,2,-2}
	Force = Derived{1,1,-2}
	Power = Derived{1,2,-3}
	Pressure = Derived{1,-1,-2}
	Frequency =	Derived{0,0,-1}
)


type Derived Dims

func (d Derived) String()string{
	switch d{
	case Energy:
		return "J"
	case Force:
		return "N"
	case Power:
		return "W"
	case Pressure:
		return "Pa"
	case Frequency:
		return "Hz"
	default:
		return Dims(d).String()
	}	
}

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



