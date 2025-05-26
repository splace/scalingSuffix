package scalingSuffix

import "fmt"

func ExampleDims(){
	fmt.Println(Time)
	fmt.Println(Dims{M:1,L:1,T:1})
	fmt.Println(Dims{M:-1,L:-1,T:-1})
	fmt.Println(Dims{M:2,T:-1})
	fmt.Println(Dims{L:1,T:-2})
	fmt.Println(Energy)
	fmt.Println(Dims(Energy))
	fmt.Println(Derived{T:-1})
	fmt.Println(DerivedsCombine(Energy,Frequency))
	// Output:
	// s
	// kg⋅m⋅s
	// kg-1⋅m-1⋅s-1
	// kg2⋅s-1
	// m⋅s-2
	// J
	// kg⋅m2⋅s-2
	// Hz
	// W
}

