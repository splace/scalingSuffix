package scalingSuffix

import "fmt"
import "os"

func ExampleFloat(){
	fmt.Println(NewFloat(1.1))
	fmt.Println(NewFloat(-1.2))
	fmt.Println(NewFloat(123.4))
	fmt.Println(NewFloat(-123.4))
	fmt.Println(NewFloat(1234.5))
	fmt.Println(NewFloat(-1234.5))
	fmt.Println(NewFloat(12345.6))
	fmt.Println(NewFloat(-12345.6))
	fmt.Println(NewFloat(1234567.8))
	fmt.Println(NewFloat(-1234567.8))
	fmt.Println(NewFloat(123456789.01))
	fmt.Println(NewFloat(-123456789.01))
	// Output:
	//1.1
	//-1.2
	//123.4
	//-123.4
	//1 234.5
	//-1 234.5
	//12 345.6
	//-12 345.6
	//1 234 567.8
	//-1 234 567.8
	//123 456 789.01
	//-123 456 789.01
}

func ExampleFloat_fmt(){
	fmt.Printf("%v\n",NewFloat(1123.123456))
	fmt.Printf("%q\n",NewFloat(1123.123456))
	// Output:
	// 1 123.123456
	// "1 123.123456"
}

func ExampleFloat_i18n(){
	langenv, exists := os.LookupEnv("LANG")
	lSep:=Sep
	if exists {
		switch langenv[:2]{
		case "en": 
			Sep=","
		default:
			Sep=" "
		}
	}
	fmt.Printf("%v\n",NewFloat(1123.123456))
	fmt.Printf("%q\n",NewFloat(1123.123456))
	Sep=lSep
	// Output:
	// 1,123.123456
	// "1,123.123456"	
}

