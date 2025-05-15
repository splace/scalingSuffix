package scalingSuffix

import "strings"
import "slices"
import "unicode/utf8"

func SpaceSep3(s string)string{
	return SepEvery(3,s," ")
}

func CommaSep3(s string)string{
	return SepEvery(3,s,",")
}

func NarrowSpaceSep3(s string)string{
	return SepEvery(3,s," ")
}

func UnderscoreSep3(s string)string{
	return SepEvery(3,s,"_")
}

func DotSep3(s string)string{
	return SepEvery(3,s,".")
}


// SepEvery adds a string between every n runes starting from the end.
func SepEvery(n uint8,s,sep string)string{
	var ss []string
	for len(s)>0{
		ns:=CutSuffixN(n,s)
		ss=append(ss,s[len(ns):])
		s=ns
	}
	slices.Reverse(ss)
	return strings.Join(ss,sep)
}

// CutSuffixN returns a string without n trailing, less if not possible, runes.
// Note: runes discarded = string[len(CutSuffixN(n,string)):]
func CutSuffixN(n uint8,s string) string{
	for range n{
		_, size := utf8.DecodeLastRuneInString(s)
		s=s[:len(s)-size]
		if len(s)==0 {
			break
		}
	}
	return s
}

