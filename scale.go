package scalingSuffix

import "strings"
import "fmt"
import "maps"

// add sep to digits above dp
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

var RSISuffices=Pivot(SISuffices)

func Scale(s string) string{
	ss:=strings.Split(s,"_")
	if len(ss)==1{
		return s
	}
	ss[1]=strings.Join(ss[:2],".")
	return TrimNumber(strings.Join(ss[1:],""))+fmt.Sprintf("%c",SISuffices[(len(ss)-1)*3])
}

//func Remove(s string) string{
//	for {
//		if s[len(s)-4:]=="_000"{
//			s=s[:len(s)-4]
//			if len(s)>4 {
//				continue
//			}
//		}
//		break
//	}
//	return s
//}

// shortens, without changing value, number containing a dp
func TrimNumber(s string) string{
	return strings.TrimSuffix(strings.TrimRight(s,"0_"),".")
}

// looks for next smaller suffix
func SmallerSuffix(e int) (int,rune){
	for {
		if r,has:=SISuffices[e];has{
			return e,r
		}
		e--
	}
}

// move second byte right by n
func MovePointRight(s string,n int) string{
	return s[0:1]+s[2:2+n]+"."+s[2+n:]
}

