package scalingSuffix

import "fmt"

func ExampleDims(){
	fmt.Println(Dims{1,1,1})
	fmt.Println(Dims{2,0,1})
	fmt.Println(Dims{0,0,2})
	// Output:
	// kg.m.s
	// kg2.s
	// s2
}

