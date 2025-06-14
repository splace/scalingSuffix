package scalingSuffix

import "strings"
import "fmt"

// apply [NumSep] to digits above dp
func Thousands(s string) string{
	b,a,f:=strings.Cut(s,".")
	b=NumSep(b)
	if f{
		return b+"."+a
	}
	return b
}

var SISuffices = map[int]rune{
	-27:'q',
	-24:'y',
	-21:'z',
	-18:'a',
	-15:'f',
	-12:'p',
	-9:'n',
	-6:'µ',
	-3:'m',
	-2:'c',
	-1:'d',
	2:'h',
	3:'k',
	6:'M',
	9:'G',
	12:'T',
	15:'P',
	18:'E',
	21:'Z',
	24:'Y',
	27:'R',
	30:'Q',
}

var SufficesSI=Pivot(SISuffices)

func Scale(s string) string{
	ss:=strings.Split(s,"_")
	if len(ss)==1{
		return s
	}
	ss[1]=strings.Join(ss[:2],".")
	return TrimNumber(strings.Join(ss[1:],""))+fmt.Sprintf("%c",SISuffices[(len(ss)-1)*3])
}

// TrimNumber removes trailing redundent, for scanning as a number, runes.
func TrimNumber(s string) string{
	return strings.TrimSuffix(strings.TrimRight(s,"0_"),".")
}

// SmallerSuffix searches for the next smaller suffix.
func SmallerSuffix(e int) (int,rune){
	for ;e>-27;e--{
		if r,has:=SISuffices[e];has{
			return e,r
		}
	}
	return e,0
}

// move second byte right by n
func MovePointRight(s string,n int) string{
	return s[0:1]+s[2:2+n]+"."+s[2+n:]
}

