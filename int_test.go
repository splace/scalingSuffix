package scalingSuffix

import "fmt"

func ExampleCutSuffixN(){
	s:="12345"
	ss:=CutSuffixN(2,s)
	fmt.Println(ss,s[len(ss):])
	// Output:
	// 123 45
}

func ExampleSepEvery3(){
	fmt.Println(SepEvery3("12"))
	fmt.Println(SepEvery3("1234"))
	fmt.Println(SepEvery3("123456"))
	fmt.Println(SepEvery3("12345678"))
	fmt.Println(SepEvery3("1234567890"))
	fmt.Println(SepEvery3("1234567890123456789"))
	fmt.Println(SepEvery3("Hello 世界"))
	// Output:
	// 12
	// 1 234
	// 123 456
	// 12 345 678
	// 1 234 567 890
	// 1 234 567 890 123 456 789
	// He llo  世界
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
	// 1 230
	// -1 230
	// 12 345
	// -12 345
	// 1 234 567
	// -1 234 567
	// 123 456 789
	// -123 456 789
}

