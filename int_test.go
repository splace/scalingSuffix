package scalingSuffix

import "fmt"

func ExampleCutSuffixN(){
	s:="12345"
	ss:=CutSuffixN(2,s)
	fmt.Println(ss,s[len(ss):])
	// Output:
	// 123 45
}

func ExampleNumSep(){
	fmt.Println(NumSep("12"))
	fmt.Println(NumSep("1234"))
	fmt.Println(NumSep("123456"))
	fmt.Println(NumSep("12345678"))
	fmt.Println(NumSep("1234567890"))
	fmt.Println(NumSep("1234567890123456789"))
	fmt.Println(NumSep("Hello 世界"))
	// Output:
	// 12
	// 1_234
	// 123_456
	// 12_345_678
	// 1_234_567_890
	// 1_234_567_890_123_456_789
	// He_llo_ 世界
}

func ExampleInt(){
	fmt.Println(NewInt(1))
	fmt.Println(NewInt(-1))
	fmt.Println(NewInt(123))
	fmt.Println(NewInt(-123))
	fmt.Println(NewInt(1230))
	fmt.Println(NewInt(-1230))
	fmt.Println(NewInt(12345))
	fmt.Println(NewInt(-12345))
	fmt.Println(NewInt(1234567))
	fmt.Println(NewInt(-1234567))
	fmt.Println(NewInt(123456789))
	fmt.Println(NewInt(-123456789))
	// Output:
	// 1
	// -1
	// 123
	// -123
	// 1_230
	// -1_230
	// 12_345
	// -12_345
	// 1_234_567
	// -1_234_567
	// 123_456_789
	// -123_456_789
}

