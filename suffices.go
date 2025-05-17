package scalingSuffix

import "fmt"
import "math"
import "golang.org/x/exp/constraints"


// suffix stores the number of scaling factors. (-ve scale down)
type Suffix int8

// calculate the suffix from the value and the scaling factor base
func SuffixFor(v,sf float64) Suffix{
	return Suffix(math.Floor(math.Log(math.Abs(v))*sf))
}

// calculate the suffices scaling from the scaling factor base
func (s Suffix) Scale(sf float64) float64{
	return math.Pow(sf,-float64(s)) // optimisation: using a -ve exponent to Pow converts a divide to a multiply.
}

const sufficesSI = "qryzafpnµm kMGTPEZYRQ"

type SISuffix struct{
	Suffix
}

func (s SISuffix) String() (_ string){
	if s.Suffix==0{return}
	return string(sufficesSI[s.Suffix+11])
}

func (s *SISuffix) Scan(state fmt.ScanState,verb rune) (err error){
	ss,_,err:=state.ReadRune()
	if err!=nil{
		return
	}
	*s=NewSISuffix(ss)
	if s.Suffix==0{
		err=fmt.Errorf("Not SI")
	}
	return
}

func NewSISuffix[I constraints.Integer](c I) (_ SISuffix){
	ci:=uint8(c)
	for i :=range len(sufficesSI){
		if sufficesSI[i]==ci{
			return SISuffix{Suffix(i-11)}
		}	
	}
	return
}


