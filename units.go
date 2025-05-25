package scalingSuffix

import "golang.org/x/exp/constraints"
import "fmt"

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



