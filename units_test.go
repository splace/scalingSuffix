package scalingSuffix

import "fmt"

func ExampleDims(){
	fmt.Println(Dims{0,0,1})
	fmt.Println(Dims{1,1,1})
	fmt.Println(Dims{-1,-1,-1})
	fmt.Println(Dims{2,0,1})
	fmt.Println(Dims{0,1,-2})
	fmt.Println(Energy)
	fmt.Println(Dims(Energy))
	fmt.Println(Derived(Time.Reciprical()))
	fmt.Println(DerivedsCombine(Energy,Derived(Time.Reciprical())))
	// Output:
	// s
	// kg⋅m⋅s
	// kg-1⋅m-1⋅s-1
	// kg2⋅s
	// m⋅s-2
	// J
	// kg⋅m2⋅s-2
	// Hz
	// W
}

