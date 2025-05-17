package scalingSuffix

import "fmt"

func ExampleNewSISuffix(){
	fmt.Printf("%q\n",NewSISuffix(uint8('1')))
	fmt.Printf("%q\n",NewSISuffix(uint8(' ')))
	fmt.Printf("%q\n",NewSISuffix(uint8('k')))
	fmt.Printf("%q\n",NewSISuffix(uint8('P')))
	// Output:
	// ""
	// "" 
	// "k"
	// "P"
}

