package scalingSuffix

import "fmt"
import "math"
import "strings"
import "golang.org/x/exp/constraints"

type value[N constraints.Float|constraints.Integer] struct{
	Value N
}

// SI is a fmt.Formatter writing values scaled by a power of 1000.
type SI[N constraints.Float|constraints.Integer] value[N]

func NewSI[N constraints.Float|constraints.Integer](v N) SI[N]{
	return SI[N]{v}
}

var sSI,fSI=1000.0,1/math.Log(1000)

func (f SI[N]) Format(fs fmt.State, v rune){
	if v=='s' {
		fmt.Fprint(fs,SISuffix{SuffixFor(float64(f.Value),fSI)})
		return
	}
	var s SISuffix
	if i:=strings.IndexByte(sufficesSI,byte(v));i>=0{
		s=SISuffix{Suffix(i-11)}
		fmat:=fmt.FormatString(fs,'v')
		n:=float64(f.Value)*s.Scale(sSI)
		if fs.Flag('+'){
			// format with dividers
			fmt.Fprintf(fs,fmat,NewFloat(n))
		}else{
			fmt.Fprintf(fs,fmat,N(n))
		}
	}else{
		// verb not found as a suffix calculate it
		s=SISuffix{SuffixFor(float64(f.Value),fSI)}
		n:=float64(f.Value)*s.Scale(sSI)
		fmt.Fprintf(fs,fmt.FormatString(fs,v),n)
	}
	// skip suffix if flagged, can be added; [ExampleSI_no_suffix]
	if fs.Flag('-'){
		return
	}
	fmt.Fprint(fs,s)	
}


const sufficesSI = "qryzafpnµm kMGTPEZYRQ"

type SISuffix struct{
	Suffix
}

func (s SISuffix) String() (_ string){
	if s.Suffix==0{return}
	return string(sufficesSI[s.Suffix+11])
}

