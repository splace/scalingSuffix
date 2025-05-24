package scalingSuffix

import "io"
import "strings"
import "slices"
import "unicode/utf8"



// NumSep adds a underscore between every 3 runes starting from the end.
func NumSep(s string)string{
	return SepEvery(3,s,"_")
}

// SepEvery adds a string between every n runes starting from the end.
func SepEvery(n int,s,sep string)string{
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
func CutSuffixN(n int,s string) string{
	for range n{
		_, size := utf8.DecodeLastRuneInString(s)
		s=s[:len(s)-size]
		if len(s)==0 {
			break
		}
	}
	return s
}

// CutPrefixN returns a string without n starting, less if not possible, runes.
// Note: runes discarded = string[:len(CutPrefixN(n,string))]
func CutPrefixN(n int,s string) string{
	for range n{
		_, size := utf8.DecodeRuneInString(s)
		s=s[size:]
		if len(s)==0 {
			break
		}
	}
	return s
}

// NumberReader Reads an embedded io.Reader, closing when reaching anything not possible for a number.
func NumberLimitedReader(r io.RuneScanner) io.Reader{
	return UntilReader(r,isNotNumber)
}

func isNotNumber(r rune) bool{
	return !strings.ContainsAny("0123456789.+-_",string(r))
}

// UntilReader Reads an embedded io.Reader, closing when a read rune casuses the provided function to return true.
func UntilReader(r io.RuneScanner, fn func(rune)bool) io.Reader{
	pr,pw:=io.Pipe() // io.PipeReader
	go func(){
		var rr rune
		var err error
		var buf [4]byte
		for {
			rr,_,err=r.ReadRune()
			if err!=nil || fn(rr){
				pw.CloseWithError(err)
				break
			}
			pw.Write(buf[:utf8.EncodeRune(buf[:],rr)])
		}
		r.UnreadRune()
	}()
	return pr
}



