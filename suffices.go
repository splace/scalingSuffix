package scalingSuffix

import "math"

// suffix stores the number of scaling factors. (-ve scale down)
type Suffix int8

// calculate the suffix from the value and the scaling factor base
func SuffixFor(v,sf float64) Suffix{
	return Suffix(math.Floor(math.Log(math.Abs(v))*sf))
}

// calculate the suffices scaling from the scaling factor base
func (s Suffix) Scale(sf float64) float64{
	return math.Pow(sf,-float64(s)) // -ve exponent to Pow removes a divide.
}

const sufficesSI = "qryzafpnµm kMGTPEZYRQ"

type SISuffix struct{
	Suffix
}

func (s SISuffix) String() (_ string){
	if s.Suffix==0{return}
	return string(sufficesSI[s.Suffix+11])
}


func NewSISuffix(c uint8) (_ SISuffix){
	for i :=range len(sufficesSI){
		if sufficesSI[i]==c{
			return SISuffix{Suffix(i-11)}
		}	
	}
	return
}


