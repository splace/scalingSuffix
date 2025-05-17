package scalingSuffix

import "fmt"

func ExampleNewSISuffix(){
	fmt.Printf("%q\n",NewSISuffix('1'))
	fmt.Printf("%q\n",NewSISuffix(' '))
	fmt.Printf("%q\n",NewSISuffix('k'))
	fmt.Printf("%q\n",NewSISuffix('P'))
	// Output:
	// ""
	// "" 
	// "k"
	// "P"
}



