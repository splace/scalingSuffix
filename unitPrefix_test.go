package scalingSuffix

import "fmt"

func ExampleSI_int(){
	var i int64=1
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+7
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+8
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+9
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+1
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewSI(i))
	// Output:
	//-1
	//12
	//-123
	//-1.23k
	//-12.3k
	//123k
	//-1.23M
	//12.3M
	//-123M
	//1.23G
	//-12.3G
	//123G
	//-1.23T
	//12.3T
	//-123T
	//1.23P
}

func ExampleSI(){
	var i float64=1
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+7
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+8
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+9
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+1
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewSI(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewSI(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewSI(i))
	// Output:
	//-1
	//12
	//-123
	//-1.23k
	//-12.3k
	//123k
	//-1.23M
	//12.3M
	//-123M
	//1.23G
	//-12.3G
	//123G
	//-1.23T
	//12.3T
	//-123T
	//1.23P
}


func ExampleIEC_int(){
	var i int64=1
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+7
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+8
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+9
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+1
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewIEC(i))
	// Output:
	//-1
	//12
	//-123
	//1.21ki
	//-12.1ki
	//121ki
	//-1.18Mi
	//11.8Mi
	//-118Mi
	//1.15Gi
	//-11.5Gi
	//115Gi
	//-1.12Ti
	//11.2Ti
	//-112Ti
	//1.1Pi
}


func ExampleIEC(){
	var i float64=1
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+7
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+8
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+9
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+1
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+2
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+3
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+4
	fmt.Printf("%.3g\n",NewIEC(i))
	i=i*10+5
	fmt.Printf("%.3g\n",NewIEC(-i))
	i=i*10+6
	fmt.Printf("%.3g\n",NewIEC(i))
	// Output:
	//-1
	//12
	//-123
	//1.21ki
	//-12.1ki
	//121ki
	//-1.18Mi
	//11.8Mi
	//-118Mi
	//1.15Gi
	//-11.5Gi
	//115Gi
	//-1.12Ti
	//11.2Ti
	//-112Ti
	//1.1Pi
}


func ExampleNewSI_suffix(){
	fmt.Printf("%s\n",NewSI(1))
	fmt.Printf("%s\n",NewSI(1000))
	fmt.Printf("%s\n",NewSI(1000000))
	fmt.Printf("%s\n",NewSI(.000001))
	// Output:
	// 
	// k
	// M
	// µ 
}

func ExampleNewIEC_suffix(){
	fmt.Printf("%s\n",NewIEC(1))
	fmt.Printf("%s\n",NewIEC(2000))
	fmt.Printf("%s\n",NewIEC(2000000))
	fmt.Printf("%s\n",NewIEC(.000002))
	// Output:
	// 
	// ki
	// Mi
	// µi 
}

func ExampleSI_no_suffix(){
	fmt.Printf("%-.3gk±0.1%[1]sB\n",NewSI(1234))
	// Output:
	// 1.23k±0.1kB
}

func ExampleSI_fixed_suffix(){
	fmt.Printf("%MB\n",NewSI(1000000000))
	fmt.Printf("%+MB\n",NewSI(1000000000))
	// Output:
	// 1000MB
	// 1_000MB
}

// use a numbers calculated scaling as forced scaling on others
func ExampleSI_suffix_reuse(){
	r:=fmt.Sprintf("%s",NewSI(1000))
	fmt.Printf("%+"+r+"B\n",NewSI(12340000))
	fmt.Printf("%+"+r+"B\n",NewSI(1234000))
	fmt.Printf("%+"+r+"B\n",NewSI(1234000000))
	// Output:
	// 12_340kB
	// 1_234kB
	// 1_234_000kB
}


func ExampleSI_suffix_reuse_linedup(){
	r:=fmt.Sprintf("%s",NewSI(1000))
	fmt.Printf("\"%+12"+r+"B\"\n",NewSI(12340000))
	fmt.Printf("\"%+12"+r+"B\"\n",NewSI(1234000))
	fmt.Printf("\"%+12"+r+"B\"\n",NewSI(1234000000))
	// Output:
	// "      12_340kB"
	// "       1_234kB"
	// "   1_234_000kB"
}

