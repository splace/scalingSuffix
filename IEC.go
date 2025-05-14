package scalingSuffix

import "fmt"
import "math"
import "strings"
import "golang.org/x/exp/constraints"

// IEC is a fmt.Formatter writing values scaled by a power of 1024.
type IEC[N constraints.Float|constraints.Integer] value[N]

func NewIEC[N constraints.Float|constraints.Integer](v N) IEC[N]{
	return IEC[N]{v}
}

var sIEC,fIEC=1024.0,1/math.Log(1024)

func (f IEC[N]) Format(fs fmt.State, v rune){
	if v=='s' {
		fmt.Fprint(fs,IECSuffix{SuffixFor(float64(f.Value),fIEC)})
		return
	}
	var s IECSuffix
	if i:=strings.IndexByte(sufficesSI,byte(v));i>=0{
		fmat:=fmt.FormatString(fs,'v')
		n:=float64(f.Value)*IECSuffix{Suffix(i-11)}.Scale(sIEC)
		if fs.Flag('+'){
			// format with dividers
			fmt.Fprintf(fs,fmat,NewFloat(n))
		}else{
			fmt.Fprintf(fs,fmat,N(n))
		}
	}else{
		// verb not found as a suffix, calculate it
		s=IECSuffix{SuffixFor(float64(f.Value),fIEC)}
		n:=float64(f.Value)*s.Scale(sIEC)
		fmat:=fmt.FormatString(fs,v)
		fmt.Fprintf(fs,fmat,n)
	}
	if fs.Flag('-'){
		return
	}
	fmt.Fprint(fs,s)
}

// a fmt.Stringer that adds the suffix 'i' to the SI suffix
type IECSuffix struct{
	Suffix
}

func (s IECSuffix) String() (_ string){
	if s.Suffix==0{
		return
	}	
	return SISuffix{s.Suffix}.String()+"i"
}

