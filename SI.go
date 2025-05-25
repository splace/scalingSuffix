package scalingSuffix

import "golang.org/x/exp/constraints"
import "fmt"
import "io"
import "bytes"
import "slices"

// SI's are fmt.Stringer and fmt.Scanner where the value is scaled using powers of 10, with metric prefixes added.
type SI[N constraints.Float | constraints.Integer] struct{value N}

func NewSI[N constraints.Float | constraints.Integer](v N) SI[N]{
	return SI[N]{v}
}

// returns number with the maximum number of 1000's removed and the suffix added.
// note: doesn't use small value SI suffix only thousands.
func (i SI[N]) String()string{
	return Scale(Thousands(fmt.Sprint(i.value)))
}

// scans from any SI suffix not just thousands.
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
	if dps,isSI:=RSISuffices[sis];isSI{
		_,err=fmt.Sscan(string(ShiftDP(bs,dps)),&s.value)
		return
	}
	_,err=fmt.Sscan(string(bs),&s.value)
	state.UnreadRune()
	return
}

func ShiftDP(bs []byte, n int)[]byte{
	if n<0 {
		return ShiftDPLeft(bs,-n)
	}	
	return ShiftDPRight(bs,n)
}

// shifts the decimal point
func ShiftDPRight(bs []byte, n int)[]byte{
	dpi:=slices.Index(bs,'.')
	if dpi<0{
		return append(bs,bytes.Repeat([]byte{'0'},n)...)
	}
	p:=bs[dpi+1:]
	if len(bs)-dpi-1<n{
		bs= append(bs[:dpi],bs[dpi+1:]...)
		return append(bs,bytes.Repeat([]byte{'0'},n-len(p))...)
	}
	bs=append(bs[:dpi],bs[dpi+1:dpi+1+n]...)
	bs=append(bs,[]byte{'.'}...)
	return append(bs,p[n:]...)
}

// shifts the decimal point
func ShiftDPLeft(bs []byte, n int)[]byte{
	dpi:=slices.Index(bs,'.')
	if dpi<0{
		if len(bs)>n{
			bs:=append(bs,bs[len(bs)-1])
			copy(bs[len(bs)-n:],bs[len(bs)-n-1:])
			bs[len(bs)-n-1]='.'
			return bs	
		}
		bs= append(bytes.Repeat([]byte{'0'},n-len(bs)),bs...)
		return append([]byte{'0','.'},bs...)
	}
	if dpi>=n{
		copy(bs[dpi-n+1:dpi+1],bs[dpi-n:])
		bs[dpi-n]='.'
		if dpi==n{
			return append([]byte{'0'},bs...)
		}
		return bs
	}
	bs= slices.Delete(bs,dpi,dpi+1)
	bs= append(slices.Repeat([]byte{'0'},n-dpi),bs...)
	return append([]byte{'0','.'},bs...)
}

//func (s *Int[N]) Scan(state fmt.ScanState,verb rune) (err error){
//	bs,err:=io.ReadAll(NumberLimitedReader(state))
//	if err!=nil{
//		return
//	}
//	sis,_,err:=state.ReadRune()
//	if err!=nil{
//		return
//	}
//	d,p,f:=strings.Cut(string(bs),".")
//	d+=strings.Repeat("0",RSISuffices[sis])
//	if f{
//		d+="."+p
//	}
//	_,err=fmt.Sscan(d,&s.Value)
//	return
//}

