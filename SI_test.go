package scalingSuffix

import "fmt"

func ExampleNewSI(){
	fmt.Println(NewSI(1_234_500))
	fmt.Println(NewSI(-12345000))
	fmt.Println(NewSI(int64(123_450_000_000_000)))
	// Output:
	// 1.2345M
	// -12.345M	
	// 123.45T
}

func ExampleSI_Scan(){
	var i SI[int64]
	fmt.Sscan("1h",&i)
	fmt.Println(i)
	fmt.Sscan("1k",&i)
	fmt.Println(i)
	fmt.Sscan("123450M",&i)
	fmt.Println(i)
	fmt.Sscan("-1000",&i)
	fmt.Println(i)
	fmt.Sscan("-1000000000000000",&i)
	fmt.Println(i)
	// Output:
	// 100
	// 1k
	// 123.45G
	// -1k
	// -1P
}


