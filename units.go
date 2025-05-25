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
	Mass = Dims{1,0,0}
	Length = Dims{0,1,0}
	Time = Dims{0,0,1}
)

var (
	Energy = Derived{1,2,-2}
	Force = Derived{1,1,-2}
	Power = Derived{1,2,-3}
	Pressure = Derived{1,-1,-2}
	Frequency =	Derived{0,0,-1}
)

type ExtendedDim uint8

const (
	ampere 	ExtendedDim = iota // A 	I
	kelvin   // K 	
	mole 	  // mol 	
	candela  // cd
)

//// electrical 
//const (
//	coulomb 	C 	electric charge 	s⋅A 	
//	volt 	V 	electric potential difference[a] 	kg⋅m2⋅s−3⋅A−1 	W/A = J/C
//	farad 	F 	capacitance 	kg−1⋅m−2⋅s4⋅A2 	C/V = C2/J
//	ohm 	Ω 	electrical resistance 	kg⋅m2⋅s−3⋅A−2 	V/A = J⋅s/C2
//	siemens 	S 	electrical conductance 	kg−1⋅m−2⋅s3⋅A2 	Ω−1
//	weber 	Wb 	magnetic flux 	kg⋅m2⋅s−2⋅A−1 	V⋅s
//	tesla 	T 	magnetic flux density 	kg⋅s−2⋅A−1 	Wb/m2
//	henry	H 	inductance 	kg⋅m2⋅s−2⋅A−2 	Wb/A
//)

//// temp
//const (
//	degree Celsius 	°C 	Celsius temperature 	K 	273.15
//)

////light
//const (
//	lumen 	lm 	luminous flux 	cd⋅sr[nc 2] 	cd⋅sr
//	lux 	lx 	illuminance 	cd⋅sr⋅m−2[nc 2] 	lm/m2
//)

//// nuclear
//const (
//	becquerel 	Bq 	activity referred to a radionuclide 	s−1 	
//	gray 	Gy 	absorbed dose, kerma 	m2⋅s−2 	J/kg
//	sievert 	Sv 	dose equivalent 	m2⋅s−2
//)

//// constsnats
//const (
//	radian[nc 1] 	rad 	plane angle 		1  2pi
//	steradian[nc 1] 	sr 	solid angle 		1  4pi
//	SpeedOfLight c	{0,1,-2} 299792458 m/s
//	PlanksContant h {1,2,-1} 6.62607015×10−34
//	BoltzmannConstant 1.380649×10−23 J⋅K−1	
//	Avagadros 6.02214076×1023 mol−1 
//	ElementaryCharge e 	1.602176634×10−19 
//)

//func Combine[Ds []D, D ~Dims](ds Ds)(cd Dims){
//	for _,d:=range ds{
//		cd.M+=d.M
//		cd.L+=d.L
//		cd.T+=d.T
//	}
//	return
//}


type Derived Dims

func (d Derived) Reciprical()Derived{
	return Derived(Dims(d).Reciprical())
}

func DerivedsCombine(ds ...Derived)(cd Derived){
	for _,d:=range ds{
		cd.M+=d.M
		cd.L+=d.L
		cd.T+=d.T
	}
	return
}

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

func DimsCombine(ds ...Dims)(cd Dims){
	for _,d:=range ds{
		cd.M+=d.M
		cd.L+=d.L
		cd.T+=d.T
	}
	return
}

func (d Dims) Reciprical()Dims{
	return Dims{-d.M,-d.L,-d.T}
}

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
		"⋅",
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



