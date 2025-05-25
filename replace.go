package scalingSuffix

import "os"

type byteReplace [2]byte

// standard byte replacment(s)
var (
	Default = byteReplace{'_',' '}
	English = byteReplace{'_',','}
	NonEnglish = []byteReplace{{'.',','},{'_','.'}}
)

// standard string modifiers
var (
	ToEnglish=ReplaceString(English)
	FromEnglish=ReplaceString(reverse(English)...)
	ToNonEnglish = ReplaceString(NonEnglish...)
	FromNonEnglish = ReplaceString(reverse(NonEnglish...)...)
)

// multi-replaces bytes in a string, ASCII only.
func ReplaceString(bps ...byteReplace) func(string)string{
	return func(s string)string{
		bs:=[]byte(s)
		replace(bs,bps...)
		return string(bs)
	}
}

// inplace byte multi-replacement, ordered.
func replace(ss []byte,bps ...byteReplace){
	for _,bp:=range bps{
		for i,s:=range ss{
			if s==bp[0]{
				ss[i]=bp[1]
			}
		}
	}
	return
}

// switch replace sense and order
func reverse(bps ...byteReplace)(r []byteReplace){
	r=make([]byteReplace,len(bps))
	for i,bp:=range bps{
		r[len(bps)-i-1]=byteReplace{bp[1],bp[0]}
	}
	return
}

// determines what byteReplace is needed for the system language
func ToSystembyteReplace(countryCode byteReplace) func(string)string{
	langenv, exists := os.LookupEnv("LANG")
	if exists{
		return I18nbyteReplace([2]byte([]byte(langenv[:2])))
	}
	return ReplaceString(Default)
}

// determines what byteReplace(s) needed for a language code
func I18nbyteReplace(countryCode [2]byte) func(string)string{
	switch countryCode{
	case byteReplace([]byte("en")): 
		return ToEnglish
	default:
		return ToNonEnglish
	}
}

