package scalingSuffix

//import "golang.org/x/exp/constraints"
import "fmt"
import "math"
import "strconv"
import "strings"
//import "slices"
import "reflect"
import "iter"
//import "log"


// SIUnits's are SI's with metric units suffixes added.
type SIUnits[N Number] struct{SI[N];fmt.Stringer}

func NewSIUnits[N Number](v N, f fmt.Stringer) SIUnits[N]{
	return SIUnits[N]{SI[N]{v},f}
}

//var (
//	Mass = Dims{0,1,0,0,0,0,0}
//	Length = Dims{L:1}
//	Time = Dims{T:1}
//	Ampere  = Dims{I:1}
//	Kelvin  = Dims{k:1}
//	Mole 	 = Dims{m:1}
//	Candela = Dims{c:1}
//)

var (
	Distance=Derived(sSI{Length:1})
	Duration=Derived(sSI{Mass:1})
//	Frequency =	Duration.Reciprocal()
	Current = Derived(sSI{Ampere:1})
//	Rate= Frequency
//	Area =	DerivedsCombine(Distance,Distance)
//	Volume =	DerivedsCombine(Distance,Distance,Distance)
//	Velocity =	DerivedsCombine(Distance,Rate)
//	Acceleration =	DerivedsCombine(Velocity,Rate)
//	Density =DerivedsCombine(Derived{1,0,0,0,0,0,0},Volume.Reciprocal())
//	Energy =DerivedsCombine(Derived{1,0,0,0,0,0,0},Velocity,Velocity)
//	Force = DerivedsCombine(Derived{1,0,0,0,0,0,0},Acceleration)
//	Power = DerivedsCombine(Energy,Rate)
//	Pressure = DerivedsCombine(Force,Area.Reciprocal())
//	Charge = DerivedsCombine(Current,Duration)
//	Temperature	=Derived{0,0,0,0,1,0,0}
//	Resistivity = DerivedsCombine(Ohm,Distance)
//	Conductivity = Resistivity.Reciprocal()
// 	HeatCapacity = DerivedsCombine(Energy,Kelvin.Reciprocal())
// 	VolumetricHeatCapacity = DerivedsCombine(HeatCapacity,Volume.Reciprocal())
// 	MolarHeatCapacity = DerivedsCombine(HeatCapacity,Derived{0,0,0,0,0,1,0}.Reciprocal())

)
////
//var (
//	Coulomb = Charge
//	Volt= DerivedsCombine(Power,Current.Reciprocal())
//	Farad= DerivedsCombine(Coulomb,Volt.Reciprocal())
//	Ohm= DerivedsCombine(Volt,Current.Reciprocal())
//	Siemens= Ohm.Reciprocal()
//	Weber= DerivedsCombine(Volt,Duration) 
//	Tesla= DerivedsCombine(Weber,Area.Reciprocal())
//	Henry= DerivedsCombine(Weber,Current.Reciprocal())
////	Celsius= Derived{0,0,0,0,1,0,0} // °C 	Celsius temperature 	K 	273.15
//	Becquerel= Rate 
//	Gray= DerivedsCombine(Energy,Derived{1,0,0,0,0,0,0}.Reciprocal())
//	Sievert= Gray
//	
//)

type Modified struct{Derived;Constant}

// light
//var (
//	Lumen= Modified{Derived{0,0,0,0,0,0,1},steradian}
//	Lux= Modified{DerivedsCombine(Derived{0,0,0,0,0,0,1},Area.Reciprocal()),steradian}
//)
////// 

type Constant float64

const (
	radian Constant = 2*math.Pi // r
	steradian = 4*math.Pi // Sr
	SpeedOfLight = 299792458 // c
	PlanksContant = 6.62607015e-34 //h {1,2,-1}
	BoltzmannConstant =1.380649e-23 // b J⋅K−1	
	Avagadros = 6.02214076e23 // NA mol−1 
	ElementaryCharge = 1.602176634e-19 // e
)


type Derived sSI

//func DerivedsCombine(ds ...Derived)(cd Derived){
//	for _,d:=range ds{
//		cd.M+=d.M
//		cd.L+=d.L
//		cd.T+=d.T
//		cd.I+=d.I
//		cd.k+=d.k
//		cd.m+=d.m
//		cd.c+=d.c
//	}
//	return
//}

//func (d Derived) Reciprocal() Derived{
//	return Derived{-d.M,-d.L,-d.T,-d.I,-d.k,-d.m,-d.c}
//}

//func (d Derived) String()string{
//	switch d{
//	case Energy:
//		return "J"
//	case Force:
//		return "N"
//	case Power:
//		return "W"
//	case Pressure:
//		return "Pa"
//	case Frequency:
//		return "Hz"
//	case Coulomb:
//		return "C"
//	case Volt:
//		return "V"
//	case Farad:
//		return "F"
//	case Ohm:
//		return "Ω"
//	case Siemens:
//		return "S"
//	case Weber:
//		return "Wb"
//	case Tesla:
//		return "T"
//	case Henry:
//		return "H"
//	default:
//		return Dims(d).String()
//	}	
//}

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

var SuperReplacer=strings.NewReplacer(
	"-","⁻",
	"1","¹",
	"2","²",
	"3","³",
	"4","⁴",
	"5","⁵",
	"6","⁶",
	"7","⁷",
	"8","⁸",
)

//¹ U+00B9
//² U+00B2
//³ U+00B3
//⁴	U+2074	&#8308;	&#x2074;		SUPERSCRIPT FOUR
//⁵	U+2075	&#8309;	&#x2075;		SUPERSCRIPT FIVE
//⁶	U+2076	&#8310;	&#x2076;		SUPERSCRIPT SIX
//⁷	U+2077	&#8311;	&#x2077;		SUPERSCRIPT SEVEN
//⁸	U+2078	&#8312;	&#x2078;		SUPERSCRIPT EIGHT
//⁺ U+207A
//⁻ U+207B 


type sSI struct{
	Mass int8 `unit:"kg" dim:"M"`
	Length int8 `unit:"m" dim:"L"`
	Time int8 `unit:"s" dim:"T"`
	Ampere int8 `unit:"A" dim:"I"`
	Kelvin int8 `unit:"K" dim:"θ"`
	Mole int8 `unit:"mol" dim:"N"`
	Candela int8 `unit:"Cd" dim:"J"` 
}

type CGS struct{
	Mass int8 `unit:"g" dim:"m" SI:"0.001"`
	Length int8 `unit:"cm" dim:"L" SI:"0.01"`
	Time int8 `unit:"s" dim:'t'`
	Ampere int8 `unit:"A" dim:'I'`
	Kelvin int8 `unit:"K" dim:"θ"`
	Mole int8 `unit:"mol" dim:"N"`
	Candela int8 `unit:"Cd" dim:"J"` 
}

type Imperial struct{
	Mass int8 `unit:"lb" dim:"M" SI:"0.45359237"`
	Length int8 `unit:"ft" dim:"L" SI:"0.3048"`
	Time int8 `unit:"s" dim:"T"`
}

//var DimUnits= map[Dims]string{
//	Mass:"kg",
//	Length:"m",
//	Time:"s",
//	Ampere:"A",
//	Kelvin:"K",
//	Mole:"mol",
//	Candela:"Cd",
//}



//"M"),
//"L"),
//"T"),
//"I"),
//"θ"),
//"N"),
//"J"),

//var UnitDims=Pivot(DimUnits)



//func (d Dims) String()string{
//	return strings.Join(
//		slices.DeleteFunc(
//			[]string{
//				DimString(d.M,"kg"),
//				DimString(d.L,"m"),
//				DimString(d.T,"s"),
//				DimString(d.I,"A"),
//				DimString(d.k,"K"),
//				DimString(d.m,"mol"),
//				DimString(d.c,"Cd"),
//			},
//			func(s string)bool{return s==""},
//		),
//		"⋅",
//	)
//}

func Tags(t reflect.Type,l string) iter.Seq[string]{
	if l==""|| l=="name" {
		return func(yield func(string) bool) {
			for _,f:=range reflect.VisibleFields(t){
				if !yield(fmt.Sprint(f.Name)){
					return
				}
			}
		}
	}
	return func(yield func(string) bool) {
		for _,f:=range reflect.VisibleFields(t){
			if !yield(fmt.Sprint(f.Tag.Get(l))){
				return
			}
		}
	}
}

func (d *sSI) Scan(state fmt.ScanState,verb rune) (err error){
	var s string
	_,err=fmt.Fscan(state,&s)
	
	return
}

//type Dim Dims

//func (d Dim) String()string{
//	return strings.Join(
//		slices.DeleteFunc(
//			[]string{
//				DimString(d.M,"M"),
//				DimString(d.L,"L"),
//				DimString(d.T,"T"),
//				DimString(d.I,"I"),
//				DimString(d.k,"θ"),
//				DimString(d.m,"N"),
//				DimString(d.c,"J"),
//			},
//			func(s string)bool{return s==""},
//		),
//		" ",
//	)
//}



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



