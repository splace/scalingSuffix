package scalingSuffix

import "os"
import "io"
import "strings"
import "fmt"

func ExampleSepEvery(){
	fmt.Println(SepEvery(1,"123","_"))
	fmt.Println(SepEvery(3,"123456789","_"))
	fmt.Println(SepEvery(3,"1","_"))
	// Output:
	// 1_2_3
	// 123_456_789
	// 1
}

func ExampleNumberUnderscoreSep(){
	fmt.Println(NumberUnderscoreSep("123"))
	fmt.Println(NumberUnderscoreSep("1234.5678900"))
	fmt.Println(NumberUnderscoreSep("1"))
	// Output:
	// 123
	// 1_234.5678900
	// 1
}

func ExampleNumberLimitedReader(){
	io.Copy(os.Stdout,NumberLimitedReader(strings.NewReader("123v")))
	fmt.Println()
	io.Copy(os.Stdout,NumberLimitedReader(strings.NewReader("123.123v")))
	fmt.Println()
	io.Copy(os.Stdout,NumberLimitedReader(strings.NewReader("-123v")))
	fmt.Println()
	io.Copy(os.Stdout,NumberLimitedReader(strings.NewReader("123_123v")))
	fmt.Println()
	// Output:
	// 123
	// 123.123
	// -123
	// 123_123
}

func ExampleShiftDP(){
	fmt.Println(string(ShiftDP([]byte("123123"),2)))
	fmt.Println(string(ShiftDP([]byte("123.123"),2)))
	fmt.Println(string(ShiftDP([]byte("123.123"),4)))
	fmt.Println(string(ShiftDP([]byte("123.123"),-2)))
	fmt.Println(string(ShiftDP([]byte("123.123"),-3)))
	fmt.Println(string(ShiftDP([]byte("123123"),-2)))
	fmt.Println(string(ShiftDP([]byte("123123"),-8)))
	fmt.Println(string(ShiftDP([]byte("123.123"),-8)))
	fmt.Println(string(ShiftDP([]byte("-123123"),2)))
	fmt.Println(string(ShiftDP([]byte("-123.123"),2)))
	fmt.Println(string(ShiftDP([]byte("-123.123"),4)))
	fmt.Println(string(ShiftDP([]byte("-123.123"),-2)))
	fmt.Println(string(ShiftDP([]byte("-123.123"),-3)))
	fmt.Println(string(ShiftDP([]byte("-123123"),-2)))
	fmt.Println(string(ShiftDP([]byte("-123123"),-8)))
	fmt.Println(string(ShiftDP([]byte("-123.123"),-8)))
	// Output:
	// 12312300
	// 12312.3
	// 1231230
	// 1.23123
	// 0.123123
	// 1231.23
	// 0.00123123
	// 0.00000123123
	// -12312300
	// -12312.3
	// -1231230
	// -1.23123
	// -0.123123
	// -1231.23
	// -0.00123123
	// -0.00000123123
}


