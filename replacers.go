package scalingSuffix

import "os"
import "strings"

var Default =strings.NewReplacer("_",",")
var English =strings.NewReplacer("_",",")
var NonEnglish =strings.NewReplacer("_",".",".",",")


func SystemReplacer(countryCode [2]byte) *strings.Replacer{
	langenv, exists := os.LookupEnv("LANG")
	if exists{
		return I18nReplacer([2]byte([]byte(langenv[:2])))
	}
	return Default
}

func I18nReplacer(countryCode [2]byte) *strings.Replacer{
	switch countryCode{
	case [2]byte([]byte("en")): 
		return English
	default:
		return NonEnglish
	}
}

