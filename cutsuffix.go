package scalingSuffix

import "strings"
import "slices"
import "unicode/utf8"

var Sep=" "

// SepEvery3 adds the string between every third rune working from the end.
func SepEvery3(s string)string{
	return SepEvery(3,s)
}

func SepEvery(n uint8,s string)string{
	var ss []string
	for len(s)>0{
		ns:=CutSuffixN(n,s)
		ss=append(ss,s[len(ns):])
		s=ns
	}
	slices.Reverse(ss)
	return strings.Join(ss,Sep)
}

// CutSuffixN returns a string without n trailing, less if not possible.
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

