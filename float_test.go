package scalingSuffix

import "fmt"

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
	//1_234.5
	//-1_234.5
	//12_345.6
	//-12_345.6
	//1_234_567.8
	//-1_234_567.8
	//123_456_789.01
	//-123_456_789.01
}

func ExampleFloat_fmt(){
	fmt.Printf("%v\n",NewFloat(1123.123456))
	fmt.Printf("%q\n",NewFloat(1123.123456))
	// Output:
	// 1_123.123456
	// "1_123.123456"
}

func ExampleFloat_i18n(){
	fmt.Printf("%v\n",NewFloat(1123.123456))
	fmt.Printf("%q\n",NewFloat(1123.123456))
	// Output:
	// 1_123.123456
	// "1_123.123456"	
}

