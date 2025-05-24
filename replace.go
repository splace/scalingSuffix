package scalingSuffix

import "os"

type replacer [2]byte

var (
	Default = replacer{'_',' '}
	English = replacer{'_',','}
	NonEnglish = []replacer{{'.',','},{'_','.'}}
	ToEnglish=ReplaceString(English)
	FromEnglish=ReplaceString(reverse(English)...)
	ToNonEnglish = ReplaceString(NonEnglish...)
	FromNonEnglish = ReplaceString(reverse(NonEnglish...)...)
)

func ReplaceString(bps ...replacer) func(string)string{
	return func(s string)string{
		bs:=[]byte(s)
		replace(bs,bps...)
		return string(bs)
	}
}

// inplace byte replacer
func replace(ss []byte,bps ...replacer){
	for _,bp:=range bps{
		for i,s:=range ss{
			if s==bp[0]{
				ss[i]=bp[1]
			}
		}
	}
	return
}

func reverse(bps ...replacer)(r []replacer){
	r=make([]replacer,len(bps))
	for i,bp:=range bps{
		r[i]=replacer{bp[1],bp[0]}
	}
	return
}

// determines what replacer is needed for the system language
func ToSystemReplacer(countryCode replacer) func(string)string{
	langenv, exists := os.LookupEnv("LANG")
	if exists{
		return I18nReplacer([2]byte([]byte(langenv[:2])))
	}
	return ReplaceString(Default)
}

// determines what replacer(s) needed for a language code
func I18nReplacer(countryCode [2]byte) func(string)string{
	switch countryCode{
	case replacer([]byte("en")): 
		return ToEnglish
	default:
		return ToNonEnglish
	}
}

