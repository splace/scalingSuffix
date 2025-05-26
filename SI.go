package scalingSuffix

import "golang.org/x/exp/constraints"
import "fmt"
import "io"
import "bytes"
import "slices" // used to generically ooerate of single bytes, avoiding slight overhead of passing byte slices.

type Number interface{
	constraints.Float | constraints.Integer
}

// SI's are fmt.Stringer and fmt.Scanner where the i/o is scaled, from the embedded value, using powers of 10 and with the metric prefix appended.
type SI[N Number] struct{value N}

func NewSI[N Number](v N) SI[N]{
	return SI[N]{v}
}

// returns the value with the maximum number of 1000's removed and replaced with the appropriate suffix added.
// note: doesn't use small, less than 1000, value SI suffixes.
func (i SI[N]) String()string{
	// TODO floats, part solved in old version
	return Scale(Thousands(fmt.Sprint(i.value)))
}

// scans from any SI suffixed number, not just SI thousands.
func (s *SI[N]) Scan(state fmt.ScanState,verb rune) (err error){
	bs,err:=io.ReadAll(NumberLimitedReader(state))
	if err!=nil{
		return
	}
	sis,_,err:=state.ReadRune()
	if err!=nil{
		if err!=io.EOF{
			return
		}
		sis=' '
	}
	if dps,isSI:=SufficesSI[sis];isSI{
		_,err=fmt.Sscan(string(ShiftDP(bs,dps)),&s.value)
		return
	}
	_,err=fmt.Sscan(string(bs),&s.value)
	state.UnreadRune()
	return
}


func ShiftDP(bs []byte, n int)[]byte{
	if bs[0]=='-'{
		return append([]byte{'-'},ShiftPosDP(bs[1:],n)...)
	}
	return ShiftPosDP(bs,n)
}

func ShiftPosDP(bs []byte, n int)[]byte{
	if n<0 {
		return ShiftDPLeft(bs,-n)
	}	
	return ShiftDPRight(bs,n)
}

// shifts the decimal point
func ShiftDPRight(bs []byte, n int)[]byte{
	dpi:=slices.Index(bs,'.')
	if dpi<0{
		// no dp so just add zeros
		return append(bs,bytes.Repeat([]byte{'0'},n)...)
	}
	p:=bs[dpi+1:]
	if len(bs)-dpi-1<n{
		// dp shifts off right end, add some zeros
		bs= append(bs[:dpi],bs[dpi+1:]...)
		return append(bs,bytes.Repeat([]byte{'0'},n-len(p))...)
	}
	// move dp inside existing
	bs=append(bs[:dpi],bs[dpi+1:dpi+1+n]...)
	bs=append(bs,"."...)
	return append(bs,p[n:]...)
}

// shifts the decimal point
func ShiftDPLeft(bs []byte, n int)[]byte{
	dpi:=slices.Index(bs,'.')
	if dpi<0{
		if len(bs)>n{
			// insert dp inside existing
			bs:=append(bs,bs[len(bs)-1])
			copy(bs[len(bs)-n:],bs[len(bs)-n-1:])
			bs[len(bs)-n-1]='.'
			return bs	
		}
		// prepend zeros and dp
		bs= append(bytes.Repeat([]byte{'0'},n-len(bs)),bs...)
		return append([]byte{'0','.'},bs...)
	}
	if dpi>=n{
		// move left runes between dp and n; overwritting old dp
		copy(bs[dpi-n+1:dpi+1],bs[dpi-n:])
		bs[dpi-n]='.'
		if dpi==n{
			// when dp now at start need prepended zero
			return append([]byte{'0'},bs...)
		}
		return bs
	}
	// dp needed is past LHS, remove dp
	bs= append(bs[:dpi],bs[dpi+1:]...)
	// prepend zeros
	bs= append(bytes.Repeat([]byte{'0'},n-dpi),bs...)
	// prepend "0."
	return append([]byte{'0','.'},bs...)
}

