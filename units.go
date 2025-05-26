package scalingSuffix

//import "golang.org/x/exp/constraints"
import "fmt"
import "math"
import "strconv"
import "strings"
import "slices"


// SIUnits's are SI's with metric units suffixes added.
type SIUnits[N Number] struct{SI[N];fmt.Stringer}

func NewSIUnits[N Number](v N, f fmt.Stringer) SIUnits[N]{
	return SIUnits[N]{SI[N]{v},f}
}

var (
	Mass = Dims{M:1}
	Length = Dims{L:1}
	Time = Dims{T:1}
	Ampere  = Dims{a:1}
	Kelvin  = Dims{k:1}
	Mole 	 = Dims{m:1}
	Candela = Dims{c:1}
)

var (
	Distance=Derived(Length)
	Duration=Derived(Time)
	Frequency =	Duration.Reciprocal()
	Current = Derived(Ampere)
	Rate= Frequency
	Area =	DerivedsCombine(Distance,Distance)
	Volume =	DerivedsCombine(Distance,Distance,Distance)
	Velocity =	DerivedsCombine(Distance,Rate)
	Acceleration =	DerivedsCombine(Velocity,Rate)
	Density =DerivedsCombine(Derived(Mass),Volume.Reciprocal())
	Energy =DerivedsCombine(Derived(Mass),Velocity,Velocity)
	Force = DerivedsCombine(Derived(Mass),Acceleration)
	Power = DerivedsCombine(Energy,Rate)
	Pressure = DerivedsCombine(Force,Area.Reciprocal())
	Coulomb = DerivedsCombine(Current,Duration)
	Volt= DerivedsCombine(Power,Current.Reciprocal()) // Derived{M:1,L:2,T:-3,a:-1} // 	V 	electric potential difference[a] 	kg⋅m2⋅s−3⋅A−1 	W/A = J/C
	Farad= DerivedsCombine(Coulomb,Volt.Reciprocal()) // Derived{M:1,LDerived{M:-1,L:-2,T:4,a:2} // F 	capacitance 	kg−1⋅m−2⋅s4⋅A2 	C/V = C2/J
	Ohm= DerivedsCombine(Volt,Current.Reciprocal()) // Derived{M:1,L Derived{M:1,L:2,T:-3,a:-2} // Ω 	electrical resistance 	kg⋅m2⋅s−3⋅A−2 	V/A = J⋅s/C2
	Siemens= Ohm.Reciprocal() // S 	electrical conductance 	kg−1⋅m−2⋅s3⋅A2 	Ω−1
	Weber= DerivedsCombine(Volt,Duration) // Derived{M:1,L:2,T:-2,a:-1} // 	Wb 	magnetic flux 	kg⋅m2⋅s−2⋅A−1 	V⋅s
	Tesla= DerivedsCombine(Weber,Area.Reciprocal()) // Derived{M:1,T:-2,a:-1} // T 	magnetic flux density 	kg⋅s−2⋅A−1 	Wb/m2
	Henry= DerivedsCombine(Weber,Current.Reciprocal()) // Derived{M:1,L:2,T:-2,a:-2} // H 	inductance 	kg⋅m2⋅s−2⋅A−2 	Wb/A
//	Celsius= Derived{0,0,0,0,1,0,0} // °C 	Celsius temperature 	K 	273.15
	Becquerel= Rate // Bq 	activity referred to a radionuclide 	s−1
	Gray= DerivedsCombine(Energy,Derived(Mass).Reciprocal()) // Derived{L:2,T:-2} // Gy 	absorbed dose, kerma 	m2⋅s−2 	J/kg
	Sievert= Gray // Sv 	dose equivalent 	m2⋅s−2
)

type Scaled struct{Derived;Constant}

// light
var (
	Lumen= Scaled{Derived(Candela),steradian} // lm 	luminous flux 	cd⋅sr[nc 2] 	cd⋅sr
	Lux= Scaled{DerivedsCombine(Derived(Candela),Area.Reciprocal()),steradian} // lx 	illuminance 	cd⋅sr⋅m−2[nc 2] 	lm/m2
)
//// 

type Constant float64

const (
	radian Constant = 2*math.Pi // r
	steradian = 4*math.Pi // Sr
	SpeedOfLight = 299792458 // c
	PlanksContant = 6.62607015e-34 //h {1,2,-1}
	BoltzmannConstant =1.380649e-23 // b J⋅K−1	
	Avagadros = 6.02214076e23 // mol−1 
	ElementaryCharge = 1.602176634e-19 // e
)

//func Combine[Ds []D, D ~Dims](ds Ds)(cd Dims){
//	for _,d:=range ds{
//		cd.M+=d.M
//		cd.L+=d.L
//		cd.T+=d.T
//	}
//	return
//}


type Derived Dims

func DerivedsCombine(ds ...Derived)(cd Derived){
	for _,d:=range ds{
		cd.M+=d.M
		cd.L+=d.L
		cd.T+=d.T
		cd.a+=d.a
		cd.k+=d.k
		cd.m+=d.m
		cd.c+=d.c
	}
	return
}
func (d Derived) Reciprocal() Derived{
	return Derived{-d.M,-d.L,-d.T,-d.a,-d.k,-d.m,-d.c}
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
	case Coulomb:
		return "C"
	case Volt:
		return "V"
	case Farad:
		return "F"
	case Ohm:
		return "Ω"
	case Siemens:
		return "S"
	case Weber:
		return "Wb"
	case Tesla:
		return "T"
	case Henry:
		return "H"
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

type Dims struct{M,L,T,a,k,m,c int8}

func DimsCombine(ds ...Dims)(cd Dims){
	for _,d:=range ds{
		cd.M+=d.M
		cd.L+=d.L
		cd.T+=d.T
	}
	return
}

func (d Dims) String()string{
	return strings.Join(
		slices.DeleteFunc(
			[]string{
				DimString(d.M,"kg"),
				DimString(d.L,"m"),
				DimString(d.T,"s"),
				DimString(d.a,"A"),
				DimString(d.k,"K"),
				DimString(d.m,"mol"),
				DimString(d.c,"Cd"),
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



